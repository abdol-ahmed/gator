package main

import (
	"github.com/abdol-ahmed/gator/internal/config"
	"github.com/abdol-ahmed/gator/internal/database"
)

type state struct {
	config *config.Config
	db     *database.Queries
}
