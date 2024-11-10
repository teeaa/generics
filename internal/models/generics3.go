package models

// type DataType interface {
// 	Address | Dob | Name
// }

type Item[T DataType] interface {
	GetBase() *Base
	SetBase(b *Base) *T
}
