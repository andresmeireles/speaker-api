package entity

type User struct {
	id        int
	name      string
	telephone string
}

func (u User) GetId() int {
	return u.id
}

func (u User) Table() string {
	return "users"
}
