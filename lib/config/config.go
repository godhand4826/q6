package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

// LoadToEnv loads flags and dotenv to environment variables.
func LoadToEnv() {
	LoadFlags()
	LoadDotEnv()
}

func LoadFlags() {
	_ = flag.String("addr", "", "server address")
	flag.Parse()

	// visit set flags
	flag.Visit(func(f *flag.Flag) {
		os.Setenv(f.Name, f.Value.String())
	})
}

func LoadDotEnv() {
	_ = godotenv.Load()
}
