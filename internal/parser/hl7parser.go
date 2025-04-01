package parser

import (
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/models"
	"strings"
)

// HL7Patient represents the patient data extracted from an HL7 message.
type HL7Patient struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
}

func ParseHL7Message(message string) (models.Patient, error) {
	// Split the HL7 message into segments using '|'
	segments := strings.Split(message, "|")

	if len(segments) < 5 {
		return models.Patient{}, errors.New("HL7 message is not properly formatted", 400)
	}

	pidSegment := segments[1]
	firstName := segments[2]
	lastName := segments[3]
	dateOfBirth := segments[4]

	patient := models.Patient{
		ID:           pidSegment,
		FirstName:    firstName,
		LastName:     lastName,
		BirthDate:    dateOfBirth,
		ResourceType: "Patient",
	}

	if patient.ID == "" {
		return models.Patient{}, errors.New("HL7 message is missing patient ID", 400)
	}

	if patient.FirstName == "" || patient.LastName == "" {
		return models.Patient{}, errors.New("HL7 message is missing patient name", 400)
	}

	return patient, nil
}
