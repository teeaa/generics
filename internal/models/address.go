package models

type Address struct {
	Base
	AddressFields
}

type AddressFields struct {
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

func (m Address) ToResponse() *AddressResponse {
	return &AddressResponse{
		Base: m.Base,
		Data: m.AddressFields,
	}
}

type AddressRequest struct {
	Data AddressFields `json:"data"`
}

func (m AddressRequest) ToModel() *Address {
	return &Address{
		Base:          Base{},
		AddressFields: m.Data,
	}
}

type AddressResponse struct {
	Base
	Data AddressFields `json:"data"`
}
