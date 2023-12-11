package entity

type Auth struct {
	Id      int
	User    User
	Hash    string
	Expired bool
}

func (a Auth) GetId() int {
	return a.Id
}

func (a Auth) Table() string {
	return "auths"
}

func (a Auth) ToJson() map[string]interface{} {
	return map[string]interface{}{
		"hash":    a.Hash,
		"user":    a.User.Id,
		"expired": a.Expired,
	}
}
