package main

import (
	"log"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/getkin/kin-openapi/openapi3"
)

func main() {
	server := &openapi3.Server{
		URL:         "http://localhost:8080",
		Description: "This is my test server",
		Variables: map[string]*openapi3.ServerVariable{
			"Variable": {
				Enum:        []interface{}{},
				Default:     "anything",
				Description: "this is server variable",
			},
		},
	}

	swagger := openapi3.Swagger{
		OpenAPI: "3.0.0",
	}

	swagger.AddServer(server)

	data, err := swagger.MarshalJSON()
	if err != nil {

		// TEMP: Snippet for debugging. remove it before commit
		spew.Printf("\n err= %+v \n\n", err)

	}

	file, err := os.Create("swagger.json")
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to create file, Error: ", err.Error())
	}

	// Write bytes to file
	byteSlice := data
	bytesWrittern, err := file.Write(byteSlice)
	if err != nil {
		log.Fatal(err)
	}

	// TEMP: Snippet for debugging. remove it before commit
	spew.Printf("\n byteswrittern= %+v \n\n", bytesWrittern)
}
