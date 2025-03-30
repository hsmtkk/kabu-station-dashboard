package misc

import (
	"log"
	"os"
	"path"
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

func MakeTodayDataDirectory() string {
	name := time.Now().Format("2006-01-02")
	dir := path.Join("./data", name)
	if _, err := os.Stat(dir); err != nil {
		os.Mkdir(dir, 0755)
	}
	return dir
}
