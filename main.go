package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chgkunm/gatorss/internal/config"
)

func main() {
	fmt.Println("----------GATORSS----------")

	gatorConfig, err := config.Read()
	var gatorState state
	gatorState.cfg = &gatorConfig
	if err != nil {
		log.Fatal(err)
	}

	cmds := commands{commandsMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatal("insufficient arguments provided")
	}

	cmd := command{commandName: os.Args[1], commandArg: os.Args[2:]}

	err = cmds.run(&gatorState, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
