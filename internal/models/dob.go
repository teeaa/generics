package models

import "time"

type Dob struct {
	Base
	Dob time.Time `json:"dob"`
}

func (m *Dob) GetBase() *Base {
	return &m.Base
}

func (m *Dob) SetBase(b *Base) {
	m.Base = *b
}
