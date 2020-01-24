package model

type Building struct {
	Id       int      `json:"id"`
	Country  string   `json:"country"`
	City     string   `json:"city"`
	Street   string   `json:"street"`
	House    string   `json:"house"`
	Location Location `json:"location"`
}
