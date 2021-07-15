
package main

import (
    "fmt"
	"github.com/xeipuuv/gojsonschema"
	"log"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	
    schemaLoader := gojsonschema.NewReferenceLoader("file:///" + filepath.Join(filepath.FromSlash(wd), "fhir.schema.json"))
    documentLoader := gojsonschema.NewReferenceLoader("file:///" + filepath.Join(filepath.FromSlash(wd), "diagnosticreport-example.json"))

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        panic(err.Error())
    }

    if result.Valid() {
        fmt.Printf("The document is valid\n")
    } else {
        fmt.Printf("The document is not valid. see errors :\n")
        for _, desc := range result.Errors() {
            fmt.Printf("- %s\n", desc)
        }
    }
}


