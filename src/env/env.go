package env

import (
	logging "atcscraper/src/log"
	"fmt"
	"os"
)

// LoadEnv Get Env Var
func LoadEnv(Key string) string {

	EnvValue := os.Getenv(Key)

	if EnvValue == "" {
		Error := fmt.Sprintf("Error Loading Env: %v", Key)
		logging.NewError(Error)
	}

	// Return Environment Variable
	return os.Getenv(Key)
}
