package utils

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

var UserConfigDir string

// Retrieves the user config location.
func GetUserConfig() (string, error) {
	UserConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	configDir := filepath.Join(UserConfigDir, "blackbox")
	return configDir, err
}

// Loads environment variables from  a .env file if it exists in the UserConfigDir
func LoadEnvVars() {
	UserConfigDir, err := GetUserConfig()
	if err != nil {
		log.Fatal(err)
	}

	switch runtime.GOOS {
	case "linux":
		err = godotenv.Load(UserConfigDir + "/.env")
	case "darwin":
		err = godotenv.Load(UserConfigDir + "/.env")
	case "windows":
		err = godotenv.Load(UserConfigDir + "\\.env")
	}
	if err != nil {
		log.Fatal("Could not read from .env:", err)
	}
}
