package main

import (
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/parser"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/parse/hl7", func(c *gin.Context) {
		var message string
		if err := c.BindJSON(&message); err != nil {
			// Handle the error using the HandleError function
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		// Parse the HL7 message
		patient, err := parser.ParseHL7Message(message)
		if err != nil {
			// Handle the error using the HandleError function
			errorResponse, statusCode := errors.HandleError(err)
			c.JSON(statusCode, errorResponse)
			return
		}

		// If the parsing succeeds, return the patient data
		c.JSON(200, patient)
	})

	r.Run(":8080")
}
