package models

type Firm struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	BuildingId uint   `json:"building_id"`
	//Building Building `json:"building"`
}

type FirmFilter struct {
	BuildingId uint `json:"building_id"`
	RubricId   uint `json:"rubric_id"`
	InRadius   uint `json:"in_radius"`
}
