package controllers

import (
	"net/http"
	"stock-trades-api/config"
	"stock-trades-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTrade(c *gin.Context) {
    var trade models.Trade

    if err := c.ShouldBindJSON(&trade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    //trade.Timestamp = time.Now()

    config.DB.Create(&trade)
    c.JSON(http.StatusCreated, trade)
}

func GetTrades(c *gin.Context) {
    var trades []models.Trade
    config.DB.Find(&trades)

    c.JSON(http.StatusOK, trades)
}

func GetTradeByID(c *gin.Context) {
    id := c.Param("id")
    var trade models.Trade

    if err := config.DB.First(&trade, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Trade not found"})
        return
    }

    c.JSON(http.StatusOK, trade)
}
