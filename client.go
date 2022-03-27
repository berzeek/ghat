package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

// parse incoming message string
func (c *client) parseMessage() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			log.Printf("client failed to read message: %s", err.Error())
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		// determine user command
		// would clean this up in later iteration
		switch cmd {
		case "/name":
			c.commands <- command{
				id:     CmdName,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CmdJoin,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CmdRooms,
				client: c,
				args:   args,
			}
		case "/all":
			c.commands <- command{
				id:     CmdAll,
				client: c,
				args:   args,
			}
		case "/exit":
			c.commands <- command{
				id:     CmdExit,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("command not recognized: %s", cmd))
		}
	}
}

// write error output
// back to terminal
func (c *client) err(err error) {
	_, err = c.conn.Write([]byte("Error: " + err.Error() + "\n"))
	if err != nil {
		return
	}
}

// write message output
// back to terminal
func (c *client) msg(msg string) {
	t := time.Now()

	// output the timestamp and message string
	_, err := c.conn.Write([]byte(fmt.Sprintf(t.Format("15:04:05")) + ":" + msg + "\n"))
	if err != nil {
		return
	}
}
