package entity

type Auth struct {
	id   int
	user User
	hash string
}

func (a Auth) GetId() int {
	return a.id
}

func (a Auth) Table() string {
	return "auths"
}
