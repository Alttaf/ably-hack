package env

import (
	"fmt"
	"os"
)

func RetrieveValue(envKey string) (string, error) {
	envValue, present := os.LookupEnv(envKey)
	if !present {
		err := fmt.Errorf("%v unset. This must be set from an .env file", envKey)
		return "", err
	}

	return envValue, nil
}
