// models/game_request.go
package models

import (
	"gorm.io/gorm"
)

// Ensure this struct is not redeclared elsewhere in the project
type GameRequest struct {
	gorm.Model
	UserID uint              `gorm:"not null" json:"userId"`
	GameID uint              `gorm:"not null" json:"gameId"`
	Status GameRequestStatus `gorm:"type:varchar(20);not null" json:"status"`

	// Associations
	User User `gorm:"foreignKey:UserID" json:"user"`
	Game Game `gorm:"foreignKey:GameID" json:"game"`
}
