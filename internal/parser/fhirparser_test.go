package parser

import (
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFHIRMessage_ValidMessage(t *testing.T) {
	// Example of a valid FHIR message (JSON format)
	fhirMessage := `{
		"id": "12345",
		"name": [{"family": "Shala", "given": ["Erza"]}],
		"gender": "female",
		"birthDate": "2004-05-06",
		"address": [{"line": ["123 Main St"], "city": "Springfield", "state": "IL", "postalCode": "62701", "country": "USA"}],
		"resourceType": "Patient"
	}`

	// Call the function to parse the FHIR message
	patient, err := ParseFHIRMessage(fhirMessage)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Define the expected patient structure
	expectedPatient := models.Patient{
		ID:           "12345",
		FirstName:    "Erza",
		LastName:     "Shala",
		Gender:       "female",
		BirthDate:    "2004-05-06",
		Address:      "123 Main St Springfield, IL 62701",
		ResourceType: "Patient",
	}

	// Compare the expected and actual patient structs
	assert.Equal(t, expectedPatient, patient)
}

// Missing ID
func TestParseFHIRMessage_InvalidMessage(t *testing.T) {
	fhirMessage := `{
		"name": [{"family": "Shala", "given": ["Erza"]}],
		"gender": "female",
		"birthDate": "2004-05-06",
		"address": [{"line": ["123 Main St"], "city": "Springfield", "state": "IL", "postalCode": "62701", "country": "USA"}],
		"resourceType": "Patient"
	}`

	parsedPatient, err := ParseFHIRMessage(fhirMessage)

	// Expecting an error because the ID field is missing
	assert.Equal(t, errors.New("FHIR message is missing required ID field", 400), err)
	assert.Equal(t, models.Patient{}, parsedPatient)
}

// invalid JSON
func TestParseFHIRMessage_InvalidJSON(t *testing.T) {
	fhirMessage := `{
		"id": "12345",
		"name": [{"family": "Shala", "given": ["Erza"]}],
		"gender": "female",
		"birthDate": "2004-05-06",
		"address": [{"line": ["123 Main St"], "city": "Springfield", "state": "IL", "postalCode": "62701", "country": "USA"}],
		"resourceType": "Patient"
	` // Missing closing bracket for the JSON

	parsedPatient, err := ParseFHIRMessage(fhirMessage)

	// Expecting an error due to malformed JSON
	assert.NotNil(t, err)
	assert.Equal(t, models.Patient{}, parsedPatient)
}

func TestParseFHIRMessage_MissingAddress(t *testing.T) {
	// FHIR message with missing address line
	fhirMessage := `{
		"id": "12345",
		"name": [{"family": "Shala", "given": ["Erza"]}],
		"gender": "female",
		"birthDate": "2004-05-06",
		"address": [{"city": "Springfield", "state": "IL", "postalCode": "62701", "country": "USA"}],
		"resourceType": "Patient"
	}`

	// Call the function to parse the FHIR message
	patient, err := ParseFHIRMessage(fhirMessage)

	// Assert that no error occurred
	assert.NoError(t, err)

	// Define the expected patient structure (address should be formatted without a line)
	expectedPatient := models.Patient{
		ID:           "12345",
		FirstName:    "Erza",
		LastName:     "Shala",
		Gender:       "female",
		BirthDate:    "2004-05-06",
		Address:      "Springfield, IL 62701",
		ResourceType: "Patient",
	}

	// Compare the expected and actual patient structs
	assert.Equal(t, expectedPatient, patient)
}
