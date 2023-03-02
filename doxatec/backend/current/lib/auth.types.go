package lib

type AuthUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Role     string `json:"role"`
	Username string `json:"username"`
	JWT      string `json:"jwt"`
}

type AuthError struct {
	IsError bool   `json:"isError"`
	Message string `json:"message"`
}
