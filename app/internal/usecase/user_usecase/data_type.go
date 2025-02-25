package user_usecase

type SignInInput struct {
	Email    string
	Password string
	IpAddr   string
}

type SignInOutput struct {
	UserId       int
	AccessToken  string
	RefreshToken string
}
