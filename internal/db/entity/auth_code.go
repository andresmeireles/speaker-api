package entity

type AuthCode struct {
	Id        int
	User      User
	Code      string
	ExpiresAt int
}

func (ac AuthCode) GetId() int {
	return ac.Id
}

func (ac AuthCode) Table() string {
	return "auth_code"
}

func (ac AuthCode) ToJson() map[string]any {
	return map[string]any{
		"user":       ac.User.GetId(),
		"code":       ac.Code,
		"expires_at": ac.ExpiresAt,
	}
}
