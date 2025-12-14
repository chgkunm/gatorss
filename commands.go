package main

import (
	"fmt"
)

type command struct {
	commandName string
	commandArg  []string
}

type commands struct {
	commandsMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler, ok := c.commandsMap[cmd.commandName]
	if !ok {
		return fmt.Errorf("unknown command")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandsMap[name] = f
}
