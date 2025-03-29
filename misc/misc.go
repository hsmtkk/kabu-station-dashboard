package misc

import (
	"log"
	"os"
	"time"
)

func RequiredEnvVar(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatalf("env var %s is not defined", name)
	}
	return val
}

const INTERVAL_MILLI_SECONDS = 500

func Interval() {
	time.Sleep(INTERVAL_MILLI_SECONDS * time.Millisecond)
}
