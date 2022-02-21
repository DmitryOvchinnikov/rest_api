package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Scheme   string
}

func (c PostgresConfig) Dialect() string {
	return "postgres"
}

func (c PostgresConfig) DSN() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

type Config struct {
	Port     int
	Env      string
	Database PostgresConfig
}

func (c Config) IsProd() bool {
	return c.Env == "prod"
}

func Load() Config {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	var c = Config{
		port,
		os.Getenv("ENV"),
		PostgresConfig{
			os.Getenv("DB_HOST"),
			dbPort,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SCHEME"),
		},
	}
	fmt.Println("Successfully loaded config")

	return c
}
