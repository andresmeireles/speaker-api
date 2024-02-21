package person

import "database/sql"

type Person struct {
	Id       int            `json:"id"`
	Name     string         `json:"name"`
	LastName sql.NullString `json:"lastName"`
	Gender   string         `json:"gender"`
}

func (p Person) GetId() int {
	return p.Id
}

func (p Person) Table() string {
	return "persons"
}

func (p Person) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"lastName": p.LastName,
		"gender":   p.Gender,
	}
}
