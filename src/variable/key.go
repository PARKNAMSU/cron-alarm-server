package variable

import "os"

var (
	USER_TOKEN_KEY            = os.Getenv("JWT_USER_TOKEN_KEY")
	REFRESH_TOKEN_ENCRYPT_KEY = os.Getenv("REFRESH_TOKEN_ENCRYPT_KEY")
	USER_PASSWORD_ENCRYPT_KEY = os.Getenv("USER_PASSWORD_ENCRYPT_KEY")
)
