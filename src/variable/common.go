package variable

import "os"

var (
	ENVIRONMENT = os.Getenv("ENVIRONMENT")
)
