package config

import "os"

var (
	ComplexApi  = os.Getenv("COMPLEX_API")
	ExternalApi = os.Getenv("EXTERNAL_API")
)
