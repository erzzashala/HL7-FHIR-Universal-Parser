package models

// Patient structure for HL7/FHIR data
type Patient struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Gender       string `json:"gender"`
	BirthDate    string `json:"birthDate"`
	Address      string `json:"address"`
	ResourceType string `json:"resourceType"`
}
