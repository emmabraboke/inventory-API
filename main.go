package main

import (
	"inventory/cmd"
)

func main() {
	// @title           Inventory API
	// @version         2.0
	// @description     Inventory App

	// @host      inventory-staging-api.onrender.com
	// @schemes http https
	// @BasePath  /api/v1

	// @securityDefinitions.apiKey  ApiKeyAuth
	// @in header
	// @name Authorization
	cmd.Setup()
}
