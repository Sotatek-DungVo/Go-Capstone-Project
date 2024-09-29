// models/required_skill.go
package models

import (
	"gorm.io/gorm"
)

type RequiredSkill struct {
	gorm.Model
	Name string `gorm:"unique;not null" json:"name"`

	// Associations
	Games []Game `gorm:"many2many:game_required_skills;" json:"games,omitempty"`
}
