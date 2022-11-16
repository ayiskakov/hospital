package entity

type Country struct {
	Cname      string `json:"cname"`
	Population int64  `json:"population,omitempty"`
}
