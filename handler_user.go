package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SuperJake03/gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	userName := cmd.Args[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

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
		return fmt.Errorf("couldn't create user: %w", err)
	}

	if err := s.cfg.SetUser(userName); err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User was created!")
	log.Printf("User created: %v+", user)
	return nil
}

func handlerReset(s *state, cmd command) error {
	if err := s.db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}

	fmt.Println("Database reset successfully!")
	return nil
}
