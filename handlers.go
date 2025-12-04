package main

import (
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArg) != 1 {
		return fmt.Errorf("usage: %s", cmd.commandName)
	}
	err := s.cfg.SetUser(cmd.commandArg[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("user switched succesfully")
	return nil
}
