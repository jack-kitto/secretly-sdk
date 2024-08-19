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

func (p *Project) Print() {
	if p == nil {
		panic("Missing param p Project")
	}
	res, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

