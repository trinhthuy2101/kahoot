package user

type BaseUserRequest struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type BaseUserResponse struct {
	// ID    int    `json:"id"`
	Token string `json:"token"`
}
