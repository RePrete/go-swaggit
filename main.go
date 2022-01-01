package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	conf := GetConf("config.yaml")
	//openapi := spec.SwaggerProps{
	//	Swagger: `2.0`,
	//	Info: &spec.Info{
	//		InfoProps: spec.InfoProps{
	//			Title:   `Example title`,
	//			Version: `1.0.0`,
	//		},
	//	},
	//	Host:     `petstore.swagger.io`,
	//	Schemes:  []string{`https`},
	//	BasePath: `/v2`,
	//	Paths: &spec.Paths{
	//		Paths: map[string]spec.PathItem{
	//			`/pet/findByStatus?status=available`: spec.PathItem{
	//				PathItemProps: spec.PathItemProps{
	//					Get: &spec.Operation{
	//						OperationProps: spec.OperationProps{
	//							Responses: &spec.Responses{
	//								ResponsesProps: spec.ResponsesProps{
	//									StatusCodeResponses: map[int]spec.Response{
	//										200: {
	//											ResponseProps: spec.ResponseProps{
	//												Description: `æsdasdadadada`,
	//											},
	//										},
	//									},
	//								},
	//							},
	//						},
	//					},
	//				},
	//			},
	//		},
	//	},
	//}

	conf.CompleteResponse()
	result, _ := json.MarshalIndent(conf, ``, `  `)
	writeToFIle(result)
}

func writeToFIle(result []byte) {
	f, err := os.OpenFile("./swagger.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(result)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
