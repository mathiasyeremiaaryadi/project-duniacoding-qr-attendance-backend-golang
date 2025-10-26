package main

import (
	"qr-attendance-backend/config"
	"qr-attendance-backend/databases"
	"qr-attendance-backend/routes"
)

func main() {
	config.LoadEnv()

	databases.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3030"))
}
