package env

import (
	"fmt"
	"os"
)

func RetrieveValue(envKey string) string {
	envValue, present := os.LookupEnv(envKey)
	if !present {
		errorString := fmt.Errorf("%v unset. This must be set from an .env file", envKey)
		fmt.Println(errorString)
		return ""
	}

	return envValue
}
