package models

import (
	"encoding/json"
	"errors"
	"time"
)
type UnixTime struct {
    time.Time
}


func (u *UnixTime) UnmarshalJSON(b []byte) error {
    var timestamp int64
    if err := json.Unmarshal(b, &timestamp); err != nil {
        return errors.New("invalid timestamp")
    }
    u.Time = time.Unix(0, timestamp*int64(time.Millisecond))
    return nil
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
    return json.Marshal(u.Time.UnixNano() / int64(time.Millisecond))
}


type Trade struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Type      string    `json:"type" binding:"required"`
    UserID    uint      `json:"user_id" binding:"required"`
    Symbol    string    `json:"symbol" binding:"required"`
    Shares    uint      `json:"shares" binding:"required,min=1,max=100"`
    Price     float64   `json:"price" binding:"required"`
    Timestamp UnixTime `json:"timestamp"`
}
