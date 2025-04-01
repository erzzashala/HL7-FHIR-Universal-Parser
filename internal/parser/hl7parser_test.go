package parser

import (
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseHL7Message_ValidMessage(t *testing.T) {
	message := "MSH|12345|Erza|Shala|2004-06-06|"

	patient, err := ParseHL7Message(message)

	assert.NoError(t, err)

	expectedPatient := models.Patient{
		ID:           "12345",
		FirstName:    "Erza",
		LastName:     "Shala",
		BirthDate:    "2004-06-06",
		ResourceType: "Patient",
	}

	assert.Equal(t, expectedPatient, patient)
}

// Malformed HL7 message
func TestParseHL7Message_MalformedMessage(t *testing.T) {
	// Define a malformed HL7 message (incorrect structure, missing fields)
	malformedMessage := "MSH|12345|Erza|Shala"

	_, err := ParseHL7Message(malformedMessage)

	// Assert that the error is not nil for malformed messages
	assert.Error(t, err)

	// Assert that the error message matches the expected message
	if customErr, ok := err.(*errors.CustomError); ok {
		assert.Equal(t, "HL7 message is not properly formatted", customErr.Message)
		assert.Equal(t, 400, customErr.StatusCode)
	}
}

func TestParseHL7Message_InvalidMessage(t *testing.T) {
	// Define an invalid HL7 message (missing required fields)
	invalidMessage := "PID|12345|||"

	_, err := ParseHL7Message(invalidMessage)

	assert.Error(t, err)

	// Assert that the error message contains expected information
	if customErr, ok := err.(*errors.CustomError); ok {
		assert.Equal(t, "HL7 message is missing patient name", customErr.Message)
		assert.Equal(t, 400, customErr.StatusCode)
	}
}
