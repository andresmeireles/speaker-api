package db

type Entity interface {
	GetId() int

	Table() string

	ToJson() map[string]interface{}
}
