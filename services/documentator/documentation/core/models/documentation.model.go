package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

type Documentation struct {
	Id              uuid.UUID          `yaml:"id,omitempty" json:"id" bson:"_id,omitempty"`
	ServiceName     string             `yaml:"service_name" json:"service_name" bson:"service_name"`
	Version         int                `yaml:"version" json:"version" bson:"version"`
	Creator         string             `yaml:"creator" json:"creator" bson:"creator"`
	Maintainer      string             `yaml:"maintainer" json:"maintainer" bson:"maintainer"`
	Git             string             `yaml:"git" json:"git" bson:"git"`
	DependsOn       DependsOn          `yaml:"depends_on,omitempty" json:"depends_on,omitempty" bson:"depends_on,omitempty"`
	Active          bool               `yaml:"active" json:"active" bson:"active"`
	Infraestructure []*Infraestructure `yaml:"infraestructure" json:"infraestructure" bson:"infraestructure"`
	Team            string             `yaml:"team,omitempty" json:"team,omitempty" bson:"team,omitempty"`
	Language        string             `yaml:"language" json:"language" bson:"language"`
	Description     string             `yaml:"description" json:"description" bson:"description"`
	CreationDate    time.Time          `yaml:"creation_date" json:"creation_date" bson:"creation_date"`
	LastUpdated     time.Time          `yaml:"last_updated" json:"last_updated" bson:"last_updated"`
	Apis            []*Api             `yaml:"apis,omitempty" json:"apis,omitempty" bson:"apis,omitempty"`
}

type DependsOn struct {
	ServiceName string    `yaml:"service_name" json:"service_name" bson:"service_name"`
	Description string    `yaml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
	ExtraInfo   []*string `yaml:"extra_info,omitempty" json:"extra_info,omitempty" json:"extra_info,omitempty"`
}

type Infraestructure struct {
	Name        string `yaml:"name" json:"name" bson:"name"`
	Description string `yaml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
}

type Api struct {
	Type        string    `yaml:"type,omitempty" json:"type,omitempty" bson:"type,omitempty"` //rest , mq
	Description string    `yaml:"description,omitempty" json:"description,omitempty" bson:"description,omitempty"`
	ExtraInfo   []*string `yaml:"extra_info,omitempty" json:"extra_info,omitempty" bson:"extra_info,omitempty"`
}

func (d *Documentation) Read(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()
	yamlFile, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(yamlFile))
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
		return err
	}
	return nil
}
