package codesender

import (
	"time"

	"github.com/andresmeireles/speaker/internal/user"
)

type AuthCode struct {
	Id        int
	UserId    int
	User      user.User `db:"-"`
	Code      string
	ExpiresAt time.Time
}

func (ac AuthCode) GetId() int {
	return ac.Id
}

func (ac AuthCode) Table() string {
	return "auth_codes"
}

func (ac AuthCode) ToJson() map[string]any {
	return map[string]any{
		"user_id":    ac.User.GetId(),
		"code":       ac.Code,
		"expires_at": ac.ExpiresAt.Format("2006-01-02 15:04:05"),
	}
}

func (a AuthCode) IsExpired() bool {
	return time.Now().After(a.ExpiresAt)
}
