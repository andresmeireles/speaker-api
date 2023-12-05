package entity

type Person struct {
	id   int
	Name string
}

func (p Person) GetId() int {
	return p.id
}

func (p Person) Table() string {
	return "persons"
}

func (p Person) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name": p.Name,
	}
}
