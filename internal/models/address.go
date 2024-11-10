package models

type Address struct {
	Base
	Street   string `json:"street"`
	City     string `json:"city"`
	PostCode string `json:"post_code"`
	Country  string `json:"country"`
}

// GetBase to retrieve base properties
func (m Address) GetBase() *Base {
	return &m.Base
}

// SetBase to set base properties
// and return the modified object
func (m Address) SetBase(b *Base) *Address {
	m.Base = *b
	return &m
}
