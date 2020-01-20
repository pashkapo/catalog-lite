package models

type Firm struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Building Building `json:"building"`
}

type Firms []Firm
