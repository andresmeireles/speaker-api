package auth

type EmailForm struct {
	Email string `json:"email"`
}

type CodeForm struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
