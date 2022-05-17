package models

import "time"

type FindDocumentation struct {
	ServiceName     string    `yaml:"service_name" json:"service_name" bson:"service_name"`
	Version         string    `yaml:"version" json:"version" bson:"version"`
	Creator         string    `yaml:"creator" json:"creator" bson:"creator"`
	Maintainer      string    `yaml:"maintainer" json:"maintainer" bson:"maintainer"`
	Git             string    `yaml:"git" json:"git" bson:"git"`
	DependsOn       string    `yaml:"depends_on,omitempty" json:"depends_on,omitempty" bson:"depends_on,omitempty"`
	Active          bool      `yaml:"active" json:"active" bson:"active"`
	Infraestructure string    `yaml:"infraestructure" json:"infraestructure" bson:"infraestructure"`
	Team            string    `yaml:"team,omitempty" json:"team,omitempty" bson:"team,omitempty"`
	Language        string    `yaml:"language" json:"language" bson:"language"`
	Description     string    `yaml:"description" json:"description" bson:"description"`
	CreationDateMax time.Time `yaml:"creation_date_max" json:"creation_date_max" bson:"creation_date_max"`
	CreationDateMin time.Time `yaml:"creation_date_min" json:"creation_date_min"`
	Apis            string    `yaml:"apis,omitempty" json:"apis,omitempty" bson:"apis,omitempty"`
}
