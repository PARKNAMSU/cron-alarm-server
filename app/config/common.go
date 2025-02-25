package config

import "os"

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)
