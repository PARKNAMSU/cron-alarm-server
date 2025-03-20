package config

import (
	"os"
	"time"
)

type REQUEST_DATA_TYPE = string

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)

var (
	JWT_ACCESS_TOKEN_PERIOD  = time.Minute * 30
	JWT_REFRESH_TOKEN_PERIOD = time.Hour * 24 * 30
	API_KEY_AVAILABLE_PERIOD = time.Hour * 24 * 30
)

var (
	REQUEST_DATA_TYPE_STRING = "string"
	REQUEST_DATA_TYPE_INT    = "int"
	REQUEST_DATA_TYPE_FLOAT  = "float"
	REQUEST_DATA_TYPE_BOOL   = "bool"
	REQUEST_DATA_TYPE_SLICE  = "slice"
)
