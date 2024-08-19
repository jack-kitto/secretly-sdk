package models

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit"
)

type Environment struct {
	ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"type:varchar(100);not null"`
	ProjectID string    `gorm:"type:uuid;not null"`
	Secrets   []Secret  `gorm:"foreignKey:EnvironmentID;constraint:OnDelete:CASCADE"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (e *Environment) Print() {
	if e == nil {
		panic("Missing param p Project")
	}
	res, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}

func Environment_fake() Environment {
	var e Environment
	gofakeit.Struct(&e)
	return e
}

func Environment_fakeMany(num int) []Environment {
	var environments []Environment
	for i := 0; i < num; i++ {
		env := Environment_fake()
		// Ensure unique IDs or any other necessary unique fields
		env.ID = gofakeit.UUID()
		env.Secrets = Secret_fakeMany(10)
		environments = append(environments, env)
	}
	return environments
}
