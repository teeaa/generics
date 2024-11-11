package models

type DataType interface {
	Address | Dob | Name
}

type Item[T DataType] interface {
	GetBase() *Base
	SetBase(b *Base) *T
}

type RequestType[T DataType] interface {
	AddressRequest | NameRequest | DobRequest
	ToModel() *T
}

type ResponseType[T DataType] interface {
	AddressResponse | NameResponse | DobResponse
}

type ResponseItem[T DataType, R ResponseType[T]] interface {
	ToResponse() *R
}
