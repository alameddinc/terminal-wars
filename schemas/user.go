package schemas

type UserRegisterPostSchema struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
