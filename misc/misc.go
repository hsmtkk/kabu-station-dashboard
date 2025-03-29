package misc

import (
	"log"
	"os"
)

func RequiredEnvVar(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("env var %s is not defined", name)
	}
	return val
}
