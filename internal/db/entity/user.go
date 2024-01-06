package entity

type User struct {
	Id    int
	Name  string
	Email string
}

func (u User) GetId() int {
	return u.Id
}

func (u User) Table() string {
	return "users"
}

func (u User) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"name":  u.Name,
		"email": u.Email,
	}
}
