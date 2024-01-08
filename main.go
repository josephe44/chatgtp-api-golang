package main

import (
	"github.com/josephe44/chatgtp-api-golang/api"
	"github.com/josephe44/chatgtp-api-golang/initializers"
	"github.com/josephe44/chatgtp-api-golang/logger"
)

func main() {

	initializers.LoadEnvVariables()

	logger.InitializeLogger("interaction.log")

	api.StartServer()
}
