package types

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthCodeSendRequest struct {
	ReceiveAccount string `json:"receiveAccount"`
	AuthType       string `json:"authType" validate:"oneof=phone email"`
}
