package variable

import "os"

var (
	USER_TOKEN_KEY = os.Getenv("JWT_USER_TOKEN_KEY")
)
