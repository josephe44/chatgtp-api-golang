package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()

	if err != nil {
		// log.Fatal("Error loading .env file")
		fmt.Print("Error loading .env file")
	}
	// Load environment variables if needed
	os.Getenv("RENDER_API_KEY")
}
