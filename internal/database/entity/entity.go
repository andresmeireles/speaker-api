package entity

type Entity interface {
	GetId() int

	Table() string
}
