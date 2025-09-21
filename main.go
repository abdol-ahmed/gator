package main

import (
	"database/sql"

	"github.com/abdol-ahmed/gator/internal/database"
	_ "github.com/lib/pq"
)

import (
	"fmt"
	"log"
	"os"

	"github.com/abdol-ahmed/gator/internal/config"
)

var programState *state

func main() {
	cfg, err := config.LoadJsonConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	dbConnection, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}

	defer dbConnection.Close()
	dbQueries := database.New(dbConnection)
	//state := &state{config: cfg,}
	programState = new(state)
	programState.config = cfg
	programState.db = dbQueries

	commands := NewCommands()
	commands.Register("login", LoginHandler)
	commands.Register("register", CreateUserHandler)
	commands.Register("reset", DeleteUsersHandler)
	commands.Register("users", GetUsersHandler)
	commands.Register("agg", RSSFeedAggregatorHandler)
	commands.Register("addfeed", CreateFeedHandler)
	commands.Register("feeds", GetFeedsHandler)

	if len(os.Args) < 2 {
		fmt.Println("too many or fewer arguments")
		fmt.Println("Usage: cli <command> [Args...]")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]

	err = commands.run(programState, command{
		Name: cmdName,
		Args: cmdArgs,
	})

	if err != nil {
		log.Fatal(err)
	}
}
