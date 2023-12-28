package auth

type AuthCode struct {
	UserId string `json:"user_id"`
	Token  string `json:"token"`
}
