package entity

type Person struct {
	Id   int
	Name string
}

func (p Person) GetId() int {
	return p.Id
}

func (p Person) Table() string {
	return "persons"
}

func (p Person) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name": p.Name,
	}
}
