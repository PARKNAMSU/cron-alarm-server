package common

type SignUpRequest struct {
	Email    string `json:"email" validate:"string"`
	Password string `json:"password" validate:"string"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"string"`
	Password string `json:"password" validate:"string"`
}

type AuthCodeSendRequest struct {
	ReceiveAccount string `json:"receiveAccount" validate:"string"`
	AuthType       string `json:"authType" validate:"oneof=phone email"`
}
