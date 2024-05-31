package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Host string `yaml:"host"`
		User string `yaml:"user"`
		Password string `yaml:"password"`
		Name string `yaml:"name"`
		Port string `yaml:"port"`
		SSLMode string `yaml:"sslmode"`
		Timezone string `yaml:"timezone"`
	} `yaml:"database"`
}

var DB * gorm.DB

func LoadConfig() (*Config, error) {
	config := &Config{}
	data, err := ioutil.ReadFile("config/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func Connect() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}
	
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		config.Database.Host,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SSLMode,
		config.Database.Timezone,
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to databases: ", err)
	} else {
		fmt.Println("Database connection successful")
	}
}