package models

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	room     *room
	commands chan<- command
}

func (c *client) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')

		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
			case "/Nick":
				c.commands <- command{
					id:     CMD_NICK,
					client: c,
					args:   args,
				}
			case "/Join":
				c.commands <- command{
					id:     CMD_JOIN,
					client: c,
					args:   args,
				}
			case "/Rooms":
				c.commands <- command{
					id:     CMD_ROOMS,
					client: c,
					args:   args,
				}
			case "/msg":
				c.commands <- command{
					id:     CMD_MSG,
					client: c,
					args:   args,
				}
			case "/quit":
			default:
			c.err(fmt.Errorf("Unknown Command %s", cmd))
		}

	}

}

func (c *client) err(e error)  {
	c.conn.Write([]byte("Err: " + e.Error() + " \n"))
}

func (c client) Msg(s string) {
	c.conn.Write([]byte("> " + s + "\n"))
}

