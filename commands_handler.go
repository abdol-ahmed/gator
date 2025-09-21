package main

import "fmt"

type command struct {
	Name string
	Args []string
}

// Commands is a struct that holds the map of command names to their handlers.
type Commands struct {
	Handlers map[string]CommandHandler
}

type CommandHandler func(*state, command) error

// NewCommands creates and initializes a new Commands struct.
func NewCommands() *Commands {
	return &Commands{
		Handlers: make(map[string]CommandHandler),
	}
}

func (c *Commands) Register(name string, f func(*state, command) error) {
	c.Handlers[name] = f
}

func (c *Commands) run(state *state, cmd command) error {
	if handler, ok := c.Handlers[cmd.Name]; ok {
		return handler(state, cmd)
	}
	return fmt.Errorf("unknown command: %s", cmd.Name)
}
