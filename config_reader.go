package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/spec"
	"io/ioutil"
	"log"
	"net/http"
)

type Swagger struct {
	*spec.Swagger
}

func GetConf(path string) *Swagger {
	s, _ := loads.Spec(path)
	return &Swagger{s.Spec()}
}

func (s *Swagger) CompleteResponse() {
	for path, pathObj := range s.Paths.Paths {
		if pathObj.Get == nil {
			continue
		}
		resp, err := http.Get(s.createUrlFromPath(path))
		statusCode := resp.StatusCode
		if err != nil {
			log.Fatalln(err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		schema := bodyToSchema(body)
		if pathObj.Get.Responses.StatusCodeResponses == nil {
			pathObj.Get.Responses.StatusCodeResponses = map[int]spec.Response{}
		}
		pathObj.Get.Responses.StatusCodeResponses[statusCode] = spec.Response{
			ResponseProps: spec.ResponseProps{
				Schema: schema,
			},
		}

	}
}

func (s *Swagger) createUrlFromPath(path string) string {
	return s.Schemes[0] + `://` + s.Host + s.BasePath + path
}

func bodyToSchema(body []byte) *spec.Schema {
	content := map[string]interface{}{}
	err := json.Unmarshal(body, &content)
	if err != nil {
		fmt.Errorf("Unmarshal error: %s", err)
	}

	return mapToSchema(content)
}

func mapToSchema(content map[string]interface{}) *spec.Schema {
	result := &spec.Schema{}
	result.Type = spec.StringOrArray{`object`}
	result.Properties = map[string]spec.Schema{}

	for key, value := range content {

		valueSchema, ok := simpleType(value)
		// If simpleType failed then is an array or a nested object
		if ok == false {
			if arrayValue, ok := value.([]interface{}); ok == true {
				if len(arrayValue) > 0 {
					if nestedSchema, ok := simpleType(arrayValue[0]); ok == true {
						valueSchema.Type = spec.StringOrArray{`array`}
						valueSchema.Items = &spec.SchemaOrArray{
							Schema: &nestedSchema,
						}
					}

					fmt.Printf("%v", arrayValue)
				}
			} else if mapValue, ok := value.(map[string]interface{}); ok == true {
				valueSchema = *mapToSchema(mapValue)
			}
		}
		result.Properties[key] = valueSchema
	}
	return result
}

func simpleType(value interface{}) (spec.Schema, bool) {
	valueSchema := spec.Schema{}
	if float64Value, ok := value.(float64); ok == true {
		valueSchema.Type = spec.StringOrArray{`number`}
		valueSchema.Example = float64Value
		return valueSchema, true
	} else if stringValue, ok := value.(string); ok == true {
		valueSchema.Type = spec.StringOrArray{`string`}
		valueSchema.Example = stringValue
		return valueSchema, true
	} else if booleanValue, ok := value.(bool); ok == true {
		valueSchema.Type = spec.StringOrArray{`string`}
		valueSchema.Example = booleanValue
		return valueSchema, true
	}
	return valueSchema, false
}
