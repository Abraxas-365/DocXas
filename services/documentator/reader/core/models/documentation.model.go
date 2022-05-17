package models

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Documentation struct {
	Version         string             `yaml:"version" json:"version"`
	Creator         string             `yaml:"creator" json:"creator"`
	Maintainer      string             `yaml:"maintainer" json:"maintainer"`
	Infraestructure []*Infraestructure `yaml:"infraestructure" json:"infraestructure"`
	Team            string             `yaml:"team,omitempty" json:"team,omitempty"`
	Language        string             `yaml:"language" json:"language"`
	Description     string             `yaml:"description" json:"description"`
	CreationDate    time.Time          `yaml:"creation_date" json:"creation_date"`
	LastUpdated     time.Time          `yaml:"last_updated" json:"last_updated"`
	Apis            []*Api             `yaml:"apis,omitempty" json:"apis,omitempty"`
}

type MainteinersInfo struct {
}

type Infraestructure struct {
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description " json:"description"`
}

type Api struct {
	Type        string      `yaml:"type" json:"type"` //rest , mq
	Description string      `yaml:"description,omitempty" json:"description,omitempty"`
	Port        string      `yaml:"port,omitempty" json:"port,omitempty"`
	Route       string      `yaml:"route,omitempty" json:"route,omitempty"`
	Data        interface{} `yaml:"data,omitempty" json:"data,omitempty"`
}

func (d *Documentation) New() {
	yamlFile, err := ioutil.ReadFile("./test.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	d.CreationDate = time.Now()
	d.LastUpdated = time.Now()
}
