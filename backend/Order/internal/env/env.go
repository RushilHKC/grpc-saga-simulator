package env

import (
	"log"
	"os"
	"strconv"
)

func GetString(key string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Environment variable not found:  %v", key)
	}

	return val
}

func GetInt(key string) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		log.Fatalf("Environment variable not found %v", key)
	}

	valAsInt, err := strconv.Atoi(val)

	if err != nil {
		log.Fatalf("Environment variable conversion failed: %s", err)
	}

	return valAsInt
}
