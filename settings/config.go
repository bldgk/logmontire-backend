package settings

import (
	"fmt"
	"github.com/tkanos/gonfig"
)

type HttpServerConfiguration struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type RabbitMQConfiguration struct {
	ConnectionString string `json:"connection_string"`
}

type DatabaseConfig struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type LocalMongoDBConfiguration struct {
	Config  DatabaseConfig `json:"config"`
	Hosts   []string       `json:"hosts"`
	Enabled bool           `json:"enabled"`
}

type AtlasMongoDBConfiguration struct {
	Config           DatabaseConfig `json:"config"`
	Hosts            []string       `json:"hosts"`
	Enabled          bool           `json:"enabled"`
	ConnectionString string         `json:"connection_string"`
}

type MongoDBConfiguration struct {
	Local LocalMongoDBConfiguration `json:"local"`
	Atlas AtlasMongoDBConfiguration `json:"atlas"`
}

type Configuration struct {
	HttpServer HttpServerConfiguration `json:"http_server"`
	RabbitMQ   RabbitMQConfiguration   `json:"rabbitmq"`
	MongoDB    MongoDBConfiguration    `json:"db"`
}

type Configurations struct {
	Settings map[string]interface{} `json:"settings"`
}

type Config interface {
	Set()
}

var (
	Settings = Configuration{}
)

func Load() {
	err := gonfig.GetConf("config.json", &Settings)
	if err != nil {
		fmt.Println(err)
	}
}

func SetSettings(config Config) {

}

func SaveSettings() {

}
