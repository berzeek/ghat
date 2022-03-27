package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

type Config struct {
	Server struct{
		IpAddress string `json:"IpAddress"`
		Port string `json:"Port"`
	} `json:"Server"`
	LogFile string `json:"LogFile"`
}

func LoadConfiguration(filename string) (Config, error) {
	var config Config
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		return config, err
	}
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	return config, err
}

func main()  {
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), "Starting Ghat application...")

	// retrieve configuration values
	config, _ := LoadConfiguration("config.json")

	// run the server routine
	s := initServer()
	go s.run()

	// bind network address
	listener, err := net.Listen("tcp", config.Server.IpAddress + ":" + config.Server.Port)
	if err != nil {
		log.Fatalf("server failed to start: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server started successfully and listening on: %s", config.Server.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("server failed to accept connection: %s", err.Error())
			continue
		}

		go s.initClient(conn)
	}
}
