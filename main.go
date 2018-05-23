package main

import (
	"fmt"
	"log"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

const (
	// Output messages
	missingArgMsg  = "The validate command requires the swagger document url to be specified"
	validSpecMsg   = "\nThe swagger spec at %q is valid against swagger specification %s\n"
	invalidSpecMsg = "\nThe swagger spec at %q is invalid against swagger specification %s.\nSee errors below:\n"
	warningSpecMsg = "\nThe swagger spec at %q showed up some valid but possibly unwanted constructs."
)

func main() {
	swaggerDoc := "swagger.json"
	specDoc, err := loads.Spec(swaggerDoc)
	if err != nil {
		log.Fatalln(err)
	}

	validate.SetContinueOnErrors(true)
	v := validate.NewSpecValidator(specDoc.Schema(), strfmt.Default)
	result, _ := v.Validate(specDoc) // returns fully detailed result with errors and warnings
	//result := validate.Spec(specDoc, strfmt.Default)		// returns single error

	if result.IsValid() {
		log.Printf(validSpecMsg, swaggerDoc, specDoc.Version())
	}
	if result.HasWarnings() {
		log.Printf(warningSpecMsg, swaggerDoc)
		log.Printf("See warnings below:\n")
		for _, desc := range result.Warnings {
			log.Printf("- WARNING: %s\n", desc.Error())
		}

	}
	if result.HasErrors() {
		str := fmt.Sprintf(invalidSpecMsg, swaggerDoc, specDoc.Version())
		for _, desc := range result.Errors {
			str += fmt.Sprintf("- %s\n", desc.Error())
		}
		log.Printf(str)
	}

}
