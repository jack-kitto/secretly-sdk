package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Project struct {
	ID           string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name         string        `gorm:"type:varchar(100);not null"`
	Environments []Environment `gorm:"foreignKey:ProjectID;constraint:OnDelete:CASCADE"`
	CreatedAt    time.Time     `gorm:"autoCreateTime"`
	UpdatedAt    time.Time     `gorm:"autoUpdateTime"`
}
