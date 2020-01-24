package model

type Firm struct {
	Id           uint      `json:"id"`
	Name         string    `json:"name"`
	PhoneNumbers []string  `json:"phone_numbers,omitempty"`
	Building     Building  `json:"building"`
	Rubrics      []*Rubric `json:"rubrics,omitempty"`
}

type FirmFilter struct {
	BuildingId uint `json:"building_id"`
	RubricId   uint `json:"rubric_id"`
	InRadius   struct {
		Radius uint     `json:"radius"`
		Point  Location `json:"point"`
	} `json:"in_radius"`
	Search string `json:"search"`
}
