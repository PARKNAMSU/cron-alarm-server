package config

import "os"

var (
	JWT_ACCESS_TOKEN_KEY      = os.Getenv("JWT_USER_TOKEN_KEY")
	JWT_REFRESH_TOKEN_KEY     = os.Getenv("JWT_REFRESH_TOKEN_KEY")
	USER_PASSWORD_ENCRYPT_KEY = os.Getenv("USER_PASSWORD_ENCRYPT_KEY")
	USER_API_ENCRYPT_KEY      = os.Getenv("USER_API_ENCRYPT_KEY")
)
