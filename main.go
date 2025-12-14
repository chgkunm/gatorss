package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/chgkunm/gatorss/internal/config"
	"github.com/chgkunm/gatorss/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	fmt.Println("----------GATORSS----------")

	gatorCfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", gatorCfg.Db_url)
	if err != nil {
		log.Fatal("error connecting to DB")
	}
	defer db.Close()
	dbQueries := database.New(db)

	gatorS := &state{
		db:  dbQueries,
		cfg: &gatorCfg,
	}

	cmds := commands{commandsMap: make(map[string]func(*state, command) error)}
	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)

	if len(os.Args) < 2 {
		log.Fatal("insufficient arguments provided")
	}

	cmd := command{commandName: os.Args[1], commandArg: os.Args[2:]}

	err = cmds.run(gatorS, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
