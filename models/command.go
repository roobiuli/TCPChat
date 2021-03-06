package models

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_JOIN commandID = 1
	CMD_MSG  commandID = 2
	CMD_QUIT commandID = 3
	CMD_ROOMS commandID = 4
)

type command struct {
	id     commandID
	client *client
	args   []string
}
