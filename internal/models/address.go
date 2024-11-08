package models

type Address struct {
	Base
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

func (m Address) GetBase() *Base {
	return &m.Base
}

func (m Address) SetBase(b *Base) *Address {
	m.Base = *b
	return &m
}
