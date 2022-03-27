package main

type commandID int

const (
	CmdName commandID = iota
	CmdJoin
	CmdRooms
	CmdAll
	CmdExit
)

type command struct {
	id commandID
	client *client
	args []string
}
