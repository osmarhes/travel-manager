package travel

import (
	"time"

	"gorm.io/gorm"
)

type TravelRequest struct {
	gorm.Model
	UserID      uint      `json:"user_id"`
	Requester   string    `json:"requester" validate:"required"`
	Destination string    `json:"destination" validate:"required"`
	Departure   time.Time `json:"departure" validate:"required"`
	Return      time.Time `json:"return" validate:"required"`
	Status      string    `json:"status" gorm:"default:solicitado" validate:"oneof=solicitado aprovado cancelado"`
}
