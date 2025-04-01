package parser

import (
	"encoding/json"
	"hl7-fhir-parser/internal/errors"
	"hl7-fhir-parser/internal/models"
	"log"
)

// ParseFHIRMessage parses a FHIR message (in JSON format)
func ParseFHIRMessage(fhirMessage string) (models.Patient, error) {
	var fhirPatient struct {
		ID   string `json:"id"`
		Name []struct {
			Family string   `json:"family"`
			Given  []string `json:"given"`
		} `json:"name"`
		Gender    string `json:"gender"`
		BirthDate string `json:"birthDate"`
		Address   []struct {
			Line    []string `json:"line"`
			City    string   `json:"city"`
			State   string   `json:"state"`
			Postal  string   `json:"postalCode"`
			Country string   `json:"country"`
		} `json:"address"`
		ResourceType string `json:"resourceType"`
	}

	var patient models.Patient

	// Decode the FHIR JSON message into the struct
	log.Println("Starting to unmarshal the FHIR message...")
	err := json.Unmarshal([]byte(fhirMessage), &fhirPatient)
	if err != nil {
		log.Printf("Error unmarshalling FHIR message: %v", err)
		return models.Patient{}, errors.New("Invalid FHIR message format", 400)
	}
	log.Println("FHIR message unmarshalled successfully")

	// Check if required fields are missing
	if fhirPatient.ID == "" {
		return models.Patient{}, errors.New("FHIR message is missing required ID field", 400)
	}
	if len(fhirPatient.Name) == 0 || fhirPatient.Name[0].Family == "" {
		return models.Patient{}, errors.New("FHIR message is missing required name field", 400)
	}

	// Map the FHIR data to the unified Patient model
	log.Printf("Mapping FHIR fields to patient struct...")

	// Handle Name fields
	patient.LastName = fhirPatient.Name[0].Family
	if len(fhirPatient.Name[0].Given) > 0 {
		patient.FirstName = fhirPatient.Name[0].Given[0]
	}

	// Handle other fields
	patient.ID = fhirPatient.ID
	patient.Gender = fhirPatient.Gender
	patient.BirthDate = fhirPatient.BirthDate
	patient.ResourceType = fhirPatient.ResourceType

	// Handle Address field
	if len(fhirPatient.Address) > 0 {
		address := fhirPatient.Address[0]
		if len(address.Line) > 0 {
			patient.Address = address.Line[0] + " " + address.City + ", " + address.State + " " + address.Postal
		} else {
			patient.Address = address.City + ", " + address.State + " " + address.Postal
		}
	}

	log.Printf("Mapped Patient: %+v", patient)
	return patient, nil
}
