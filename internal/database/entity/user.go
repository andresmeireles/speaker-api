package entity

type User struct {
	id        int
	Name      string
	Telephone string
}

func (u User) GetId() int {
	return u.id
}

func (u User) Table() string {
	return "users"
}

func (u User) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name":      u.Name,
		"telephone": u.Telephone,
	}
}
