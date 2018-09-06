package utils

import (
	"gopkg.in/yaml.v2"
	"log"
	"io/ioutil"
)

var (
	Config Configuration
)

type Configuration struct {
	Connections Connections `yaml:"database"`
	Selectors   Selectors   `yaml:"selectors"`
}

type Connections struct {
	User   string `yaml:"user"`
	Pass   string `yaml:"pass"`
	Base   string `yaml:"base"`
	Get    string `yaml:"get"`
	Update string `yaml:"update"`
}

type Selectors struct {
	MercadoLibre Marketplace `yaml:"mercadolibre"`
	Ripley       Marketplace `yaml:"ripley"`
}

type Marketplace struct {
	Price string `yaml:"price"`
}

func LoadConfig(filename string) {
	log.Println("CONFIGURATION MODE: YAML")
	if bytes, err := ioutil.ReadFile(filename); err != nil {
		panic(err)
	} else {
		if err := yaml.Unmarshal(bytes, &Config); err != nil {
			panic(err)
		}
	}
}
