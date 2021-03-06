package main

import (
	"github.com/piotrkowalczuk/netwars-backend/database"
	"encoding/xml"
	"os"
	"log"
)

const filePath = "config.xml"

type Config struct {
	Server Server `xml:"server"`
	Redis database.RedisConfig `xml:"redis"`
	Postgre database.PostgreConfig `xml:"postgre"`
}

type Server struct {
	Host string `xml:"host"`
	Port string `xml:"port"`
	ReadTimeout int `xml:"read_timeout"`
	WriteTimeout int `xml:"write_timeout"`
}

func openFile() (file *os.File) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Printf("Cannot open configuration file: %v\n", err)
		os.Exit(1)
	}

	return
}

func ReadConfiguration() (config *Config){
	config = new(Config)
	file := openFile()
	defer file.Close()

	decoder := xml.NewDecoder(file)
	decoder.Decode(&config)

	return
}


