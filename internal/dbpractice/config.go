package dbpractice

import "os"

type Config struct {
	ListenAddr  string
	DatabaseUrl string
}

func NewConfig() *Config {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		panic("DATABASE_URL not set")
	}

	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}

	return &Config{
		ListenAddr:  listenAddr,
		DatabaseUrl: dbUrl,
	}
}
