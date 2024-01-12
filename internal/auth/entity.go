package auth

import "github.com/andresmeireles/speaker/internal/user"

type Auth struct {
	Id      int
	UserId  int
	User    user.User
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
		"user_id": a.UserId,
		"expired": a.Expired,
	}
}
