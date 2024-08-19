package secretly

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/brianvoe/gofakeit"
	"github.com/lucsky/cuid"
	"gopkg.in/yaml.v3"
)

type Secret struct {
	ID            string    `gorm:"primaryKey;type:text"`
	Name          string    `gorm:"type:varchar(100);not null"`
	Value         string    `gorm:"type:text;not null"`
	ProjectID     string    `gorm:"type:text;not null"`
	EnvironmentID string    `gorm:"type:text;not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
}

func (s *Secret) Print() {
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

func (s *Secret) ToYaml(filePath string) {
	yamlFile, err := yaml.Marshal(s)
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

func Secret_build(name string, value string, project Project, environment Environment) Secret {
	return Secret{
		ID:            cuid.New(),
		Name:          name,
		Value:         value,
		ProjectID:     project.ID,
		EnvironmentID: environment.ID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
}
