package models

import "time"

type Trade struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Type      string    `json:"type" binding:"required"`
    UserID    uint      `json:"user_id" binding:"required"`
    Symbol    string    `json:"symbol" binding:"required"`
    Shares    uint      `json:"shares" binding:"required,min=1,max=100"`
    Price     float64   `json:"price" binding:"required"`
    Timestamp time.Time `json:"timestamp"`
}
