package cfg

import (
	"os"
	"strconv"
)

type Config struct {
	DBType     string
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	HTTPPort int

	JWTSecret string
}

// Load gets the needed environmental variables needed to run the server
func Load() (*Config, error) {
	cfg := &Config{}

	// Load env vars related to the DB connection
	cfg.DBType = os.Getenv("DB_TYPE")

	if cfg.DBType == "postgres" {
		cfg.DBHost = os.Getenv("DB_HOST")
		port, err := strconv.Atoi(os.Getenv("DB_PORT"))
		if err != nil {
			return nil, err
		}
		cfg.DBPort = port
		cfg.DBUser = os.Getenv("DB_USER")
		cfg.DBPassword = os.Getenv("DB_PASS")
		cfg.DBName = os.Getenv("DB_NAME")
	}

	// Load env vars related to server ports
	port, err := strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.HTTPPort = port

	// Load env vars related to jwt auth
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	return cfg, nil
}
