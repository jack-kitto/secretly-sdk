package models

import "time"

type Secret struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Value         string    `gorm:"type:text;not null"`
	ProjectID     string    `gorm:"type:uuid;not null"`
	EnvironmentID string    `gorm:"type:uuid;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}
