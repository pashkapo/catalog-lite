package models

type Building struct {
	Id       int    `json:"id"`
	Address  string `json:"address"`
	Location Point  `json:"location"`
}
