package main

import (
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/parser"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/parse/hl7", func(c *gin.Context) {
		// Read the raw HL7 message from the request body
		message, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		// Parse the HL7 message
		patient, err := parser.ParseHL7Message(string(message))
		if err != nil {
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		// Return the parsed patient data as JSON
		c.JSON(200, patient)
	})

	// Endpoint to Parse FHIR Messages**
	r.POST("/parse/fhir", func(c *gin.Context) {
		message, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		patient, err := parser.ParseFHIRMessage(string(message))
		if err != nil {
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		c.JSON(200, patient)
	})

	r.Run(":8080")
}
