package config

import (
	"os"
	"time"
)

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)

var (
	JWT_ACCESS_TOKEN_PERIOD  = time.Minute * 30
	JWT_REFRESH_TOKEN_PERIOD = time.Hour * 24 * 30
)
