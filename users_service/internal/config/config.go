package config

import "os"

var (
	DatabaseDsn = os.Getenv("DATABASE_DSN")
)
