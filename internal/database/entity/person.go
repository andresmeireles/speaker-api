package entity

type Person struct {
	id   int
	name string
}

func (p Person) GetId() int {
	return p.id
}

func (p Person) Table() string {
	return "persons"
}
