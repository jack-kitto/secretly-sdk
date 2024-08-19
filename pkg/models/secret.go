package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Secret struct {
	ID            string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Value         string    `gorm:"type:text;not null"`
	ProjectID     string    `gorm:"type:uuid;not null"`
	EnvironmentID string    `gorm:"type:uuid;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (s *Secret) Print(filePath string) {
	if s == nil {
		panic("Missing param p Project")
	}
	res, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func Secret_fake() Secret {
	var s Secret
	gofakeit.Struct(&s)
	return s
}

func Secret_fakeMany(num int) []Secret {
	var secrets []Secret
	for i := 0; i < num; i++ {
		secret := Secret_fake()
		// Ensure unique IDs or any other necessary unique fields
		secret.ID = gofakeit.UUID()
		secrets = append(secrets, secret)
	}
	return secrets
}
