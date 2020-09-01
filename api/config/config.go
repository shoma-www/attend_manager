package config

import (
	"errors"
	"io/ioutil"
	"log"

	yaml "github.com/goccy/go-yaml"
)

// Error
var (
	ErrorFileLoad = errors.New("Error loading the config file")
	ErrorYamlLoad = errors.New("Error loading the yaml file")
)

type Server struct {
	Addr string
}

type Grpc struct {
	Addr string
	Port int
}

// Config サーバーコンフィグ
type Config struct {
	Server
	Client struct {
		Grpc
	}
}

// LoadConfig コンフィグをロード
func LoadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil, ErrorFileLoad
	}

	var config Config
	if err = yaml.Unmarshal(buf, &config); err != nil {
		log.Println(err)
		return nil, ErrorYamlLoad
	}

	return &config, nil
}
