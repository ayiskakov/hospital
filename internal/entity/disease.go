package entity

type DiseaseType struct {
	ID          int64  `json:"id"`
	Description string `json:"description,omitempty"`
}

type DiseaseTypeUpdate struct {
	Description *string `json:"description,omitempty"`
}

type Disease struct {
	DiseaseType DiseaseType `json:"disease_type"`
	Description string      `json:"description,omitempty"`
	DiseaseCode string      `json:"disease_code"`
	Pathogen    string      `json:"pathogen,omitempty"`
}

type DiseaseUpdate struct {
	ID          *int64  `json:"id,omitempty"`
	Description *string `json:"description,omitempty"`
	DiseaseCode *string `json:"disease_code,omitempty"`
	Pathogen    *string `json:"pathogen,omitempty"`
}
