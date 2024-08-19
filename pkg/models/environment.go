package models

import "time"

type Environment struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(100);not null"`
	ProjectID string    `gorm:"type:uuid;not null"`
	Secrets   []Secret  `gorm:"foreignKey:EnvironmentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
