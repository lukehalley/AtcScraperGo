package env

import (
	"log"
	"os"
)

// LoadEnv Get Env Var
func LoadEnv(Key string) string {

	EnvValue := os.Getenv(Key)

	if EnvValue == "" {
		log.Fatalf("Error Loading Env Var: '%s'", Key)
	}

	// Return Environment Variable
	return os.Getenv(Key)
}
