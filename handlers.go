package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/chgkunm/gatorss/internal/database"
)

func printUser(user database.User) {
	fmt.Printf(" * ID:      %v\n", user.ID)
	fmt.Printf(" * Name:    %v\n", user.Name)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.commandArg) != 1 {
		return fmt.Errorf("usage: %s", cmd.commandName)
	}

	name := cmd.commandArg[0]

	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("user switched succesfully")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.commandArg) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.commandName)
	}

	name := cmd.commandArg[0]

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("user created succesfully")
	printUser(user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.commandArg) != 0 {
		return fmt.Errorf("usage: %s", cmd.commandName)
	}
	err := s.db.DeleteAll(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't delete all users: %w", err)
	}
	fmt.Println("all users deleted")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.commandArg) != 0 {
		return fmt.Errorf("usage: %s", cmd.commandName)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error retriving users")
	}

	for _, user := range users {
		if user.Name == s.cfg.Current_user_name {
			fmt.Printf("* %s (current)\n", user.Name)
			continue
		}
		fmt.Printf("* %s\n", user.Name)
	}
	return nil
}
