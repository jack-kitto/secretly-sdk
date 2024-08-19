package models

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"gopkg.in/yaml.v3"
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

func Project_fake() Project {
	var p Project
	gofakeit.Struct(&p)
	p.Environments = Environment_fakeMany(5)
	return p
}

func (p *Project) ToYaml(filePath string) {
	yamlFile, err := yaml.Marshal(p)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = io.WriteString(f, string(yamlFile))
	if err != nil {
		panic(err)
	}
}
