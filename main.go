package main

import (
	"stock-trades-api/config"
	"stock-trades-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    config.ConnectDatabase()
    routes.InitializeRoutes(r)

    r.Run(":8080")
}


