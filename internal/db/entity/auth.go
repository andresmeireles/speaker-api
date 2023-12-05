package entity

type Auth struct {
	id   int
	User User
	Hash string
}

func (a Auth) GetId() int {
	return a.id
}

func (a Auth) Table() string {
	return "auths"
}

func (a Auth) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"hash": a.Hash,
		"user": a.User.id,
	}
}
