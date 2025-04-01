package models

type Observation struct {
	ID        string `json:"id"`         // Unique observation ID
	Code      string `json:"code"`       // Code representing the observation (e.g., lab test)
	Value     string `json:"value"`      // The result value of the observation
	PatientID string `json:"patient_id"` // Associated patient ID
}
