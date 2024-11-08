package models

type Name struct {
	Base
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
