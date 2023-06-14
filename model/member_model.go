package model

type MemberModel struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	PasswordHash string `json:"password"`
}

type MemberLoginModel struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}
