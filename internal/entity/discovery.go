package entity

import "time"

type Discovery struct {
	FirstEncDate time.Time `json:"first_enc_date,omitempty"`
	Disease      Disease   `json:"disease"`
	Country      Country   `json:"country"`
}

type DiscoveryUpdate struct {
	Cname        *string    `json:"cname,omitempty"`
	DiseaseCode  *string    `json:"disease_code,omitempty"`
	FirstEncDate *time.Time `json:"first_enc_date,omitempty"`
}
