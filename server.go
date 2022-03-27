package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net"
	"strings"
)

type server struct {
	rooms    map[string]*room
	commands chan command
}

// initialize server
func initServer() *server {
	return &server{
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

// find incoming command
// then run
func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CmdName:
			s.name(cmd.client, cmd.args)
		case CmdJoin:
			s.join(cmd.client, cmd.args)
		case CmdRooms:
			s.listRooms(cmd.client, cmd.args)
		case CmdAll:
			s.all(cmd.client, cmd.args)
		case CmdExit:
			s.exit(cmd.client, cmd.args)
		}
	}
}

// initialize client
func (s *server) initClient(conn net.Conn) {
	log.Printf("client successfully initialized and connected: %s", conn.RemoteAddr().String())

	// create random default username
	var newUser = fmt.Sprintf("Anonymous%d", rand.Intn(1000))

	c := client{
		conn:     conn,
		name:     newUser,
		commands: s.commands,
	}

	c.parseMessage()
}

// change username
func (s *server) name(c *client, args []string) {
	c.name = args[1]
	c.msg(fmt.Sprintf("Username succesfully changed to: %s", c.name))
}

// join a room
// create if none exists
func (s *server) join(c *client, args []string) {
	roomName := args[1]
	// check first if room exists
	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.exitRoom(c)

	c.room = r

	r.announce(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("Welcome to the room: #%s", r.name))
}

// list all rooms
func (s *server) listRooms(c *client, args []string) {
	var rooms []string
	for name := range s.rooms {
		rooms = append(rooms, name)
	}

	c.msg(fmt.Sprintf("List of available rooms: #%s", strings.Join(rooms, ", #")))
}

// delete the client
// from room members
func (s *server) exitRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.announce(c, fmt.Sprintf("%s has left the room", c.name))
	}
}

// sends a message
// to all clients
func (s *server) all(c *client, args []string) {
	if c.room == nil {
		c.err(errors.New("you must join a room to send a message"))
		return
	}

	c.room.announce(c, c.name+": "+strings.Join(args[1:len(args)], " "))
}

// ends the connection
func (s *server) exit(c *client, args []string) {
	log.Printf("Client disconnected: %s", c.conn.RemoteAddr().String())

	s.exitRoom(c)

	c.msg("Session ended")
	err := c.conn.Close()
	if err != nil {
		return
	}
}
