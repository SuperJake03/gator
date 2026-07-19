package main

import (
	"github.com/SuperJake03/gator/internal/config"
	"github.com/SuperJake03/gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}
