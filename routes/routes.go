package routes

import (
	"stock-trades-api/controllers"
	"stock-trades-api/utils"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    protected := router.Group("/", utils.AuthMiddleware())
    {
        protected.POST("/trades", controllers.CreateTrade)
        protected.GET("/trades", controllers.GetTrades)
        protected.GET("/trades/:id", controllers.GetTradeByID)
    }
}

