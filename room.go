package main

import (
	"log"
	"net"
	"os"
)

type room struct {
	name string
	members map[net.Addr]*client // TODO clean this up
}

// announce message
// to all connected clients
func (r *room) announce(announcer *client, msg string) {

	// retrieve logfile path from config file
	config, _ := LoadConfiguration("config.json")

	// log the message
	// then send to clients
	for addr, m := range r.members {
		if addr != announcer.conn.RemoteAddr() {
			// If the file doesn't exist, create it or append to the file
			file, err := os.OpenFile(config.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
			}

			// sets the output destination for the logger
			log.SetOutput(file)

			log.Println(msg)

			m.msg(msg)
		}
	}
}
