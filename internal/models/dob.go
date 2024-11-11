package models

import "time"

type Dob struct {
	Base
	DobFields
}

type DobFields struct {
	Dob time.Time `json:"dob"`
}

func (m Dob) GetBase() *Base {
	return &m.Base
}

func (m Dob) SetBase(b *Base) *Dob {
	m.Base = *b
	return &m
}

func (m Dob) ToResponse() *DobResponse {
	return &DobResponse{
		Base: m.Base,
		Data: m.DobFields,
	}
}

type DobRequest struct {
	Data DobFields `json:"data"`
}

func (m DobRequest) ToModel() *Dob {
	return &Dob{
		Base:      Base{},
		DobFields: m.Data,
	}
}

type DobResponse struct {
	Base
	Data DobFields `json:"data"`
}
