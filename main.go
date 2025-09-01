package main

import (
	"github.com/rafinhacuri/api-go.git/config"
	"github.com/rafinhacuri/api-go.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Error initializing the application: %v", err)
		return
	}

	router.InitializeRouter(":8080")
}
