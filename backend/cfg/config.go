package cfg

import (
	"os"
	"strconv"
)

// Config contains all the required config fields
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string

	GRPCPort int
	HTTPPort int

	JWTSecret string
}

// Load reads the required Environmental Variables
func Load() (*Config, error) {
	cfg := &Config{}

	// Load env vars related to the DB connection
	cfg.DBHost = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.DBPort = port
	cfg.DBUser = os.Getenv("DB_USER")
	cfg.DBPassword = os.Getenv("DB_PASS")
	cfg.DBName = os.Getenv("DB_NAME")

	// Load env vars related to server ports
	port, err = strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.GRPCPort = port

	port, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.HTTPPort = port

	// Load env vars related to jwt auth
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	return cfg, nil
}
