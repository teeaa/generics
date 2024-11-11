package models

type Name struct {
	Base
	NameFields
}

type NameFields struct {
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
}

func (m Name) GetBase() *Base {
	return &m.Base
}

func (m Name) SetBase(b *Base) *Name {
	m.Base = *b
	return &m
}

func (m Name) ToResponse() *NameResponse {
	return &NameResponse{
		Base: m.Base,
		Data: m.NameFields,
	}
}

type NameRequest struct {
	Data NameFields `json:"data"`
}

func (m NameRequest) ToModel() *Name {
	return &Name{
		Base:       Base{},
		NameFields: m.Data,
	}
}

type NameResponse struct {
	Base
	Data NameFields `json:"data"`
}
