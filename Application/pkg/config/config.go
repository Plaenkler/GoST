package config

import (
	"log"
	"os"

	"github.com/plaenkler/gost/pkg/handler"
	"gopkg.in/yaml.v3"
)

var instance *Config

type Config struct {
	Port      string `yaml:"port"`
	PathToDB  string `yaml:"pathToDB"`
	PathForDB string `yaml:"pathForDB"`
}

func GetConfig() *Config {
	defer handler.HandlePanic("config")

	if instance == nil {
		initConfig()
	}

	return instance
}

func initConfig() {
	defer handler.HandlePanic("config")

	instance = &Config{}

	if _, err := os.Stat("./config.yaml"); err != nil {
		createConfig()
	}

	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalf("[config] could not open file - error: %s", err.Error())
	}
	defer file.Close()

	d := yaml.NewDecoder(file)

	if err := d.Decode(&instance); err != nil {
		log.Fatalf("[config] could not decode file - error: %s", err.Error())
	}
}

func createConfig() {
	defer handler.HandlePanic("config")

	config := Config{
		Port:      ":80",
		PathToDB:  "",
		PathForDB: "gost.db",
	}

	data, err := yaml.Marshal(&config)
	if err != nil {
		log.Fatalf("[config] marshal failed - error: %s", err.Error())
	}

	err = os.WriteFile("config.yaml", data, 0644)
	if err != nil {
		log.Fatalf("[config] unable to write data - error: %s", err.Error())
	}

	log.Print("[config] Created config.yaml exiting...")
	os.Exit(0)
}
