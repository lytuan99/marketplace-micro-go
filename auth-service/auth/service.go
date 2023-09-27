package auth

type Service interface {
	Login(username, password string)
	Register()
	Logout()
	VerifyToken(accessToken string)
	GetMe()
}
