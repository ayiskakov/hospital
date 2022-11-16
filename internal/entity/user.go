package entity

type User struct {
	Email   string  `json:"email"`
	Name    string  `json:"name,omitempty"`
	Surname string  `json:"surname,omitempty"`
	Salary  int64   `json:"salary,omitempty"`
	Phone   string  `json:"phone,omitempty"`
	Country Country `json:"country,omitempty"`
}

type UserUpdate struct {
	Email   *string `json:"email,omitempty"`
	Name    *string `json:"name,omitempty"`
	Surname *string `json:"surname,omitempty"`
	Salary  *int64  `json:"salary,omitempty"`
	Phone   *string `json:"phone,omitempty"`
	Cname   *string `json:"cname,omitempty"`
}

type PublicServant struct {
	Department string `json:"department"`
	User       User   `json:"user"`
}

type PublicServantUpdate struct {
	Department *string `json:"department"`
}

type Doctor struct {
	Degree string `json:"degree"`
	User   User   `json:"user"`
}

type DoctorUpdate struct {
	Degree *string `json:"degree,omitempty"`
}

type Specialize struct {
	DiseaseType DiseaseType `json:"disease_type"`
	Doctor      Doctor      `json:"doctor"`
}
