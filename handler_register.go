package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SuperJake03/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]
	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      userName,
	}
	user, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Println("couldn't create user")
		os.Exit(1)
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return err
	}

	fmt.Println("User was created!")
	log.Printf("User created: %v+", user)
	return nil
}
