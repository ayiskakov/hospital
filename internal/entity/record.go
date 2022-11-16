package entity

type Record struct {
	PublicServant PublicServant `json:"public_servant"`
	Country       Country       `json:"country"`
	Disease       Disease       `json:"disease"`
	TotalDeaths   int64         `json:"total_deaths"`
	TotalPatients int64         `json:"total_patients"`
}

type RecordUpdate struct {
	TotalDeaths   *int64 `json:"total_deaths"`
	TotalPatients *int64 `json:"total_patients"`
}
