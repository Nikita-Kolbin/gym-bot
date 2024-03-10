package sqlite

import (
	"fmt"
	"gym-bot/internal/storage"
)

const (
	defaultInTrain = 0
	defaultState   = 0
)

func (s *Storage) CreateUser(username string) error {
	q := `INSERT INTO users (username, in_train, state) VALUES (?, ?, ?)`

	ok, err := s.UserIsExists(username)
	if err != nil {
		return fmt.Errorf("can't create user: %w", err)
	}
	if ok {
		return nil
	}

	if _, err = s.db.Exec(q, username, defaultInTrain, defaultState); err != nil {
		return fmt.Errorf("can't create user: %w", err)
	}

	if err = s.createCreateExercise(username); err != nil {
		return fmt.Errorf("can't create user: %w", err)
	}

	return nil
}

func (s *Storage) UserIsExists(username string) (bool, error) {
	q := `SELECT COUNT(*) FROM users WHERE username = ?`

	var count int
	if err := s.db.QueryRow(q, username).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check user exists: %w", err)
	}

	return count > 0, nil
}

func (s *Storage) createCreateExercise(username string) error {
	q := `INSERT INTO create_exercises (username) VALUES (?)`

	if _, err := s.db.Exec(q, username); err != nil {
		return fmt.Errorf("can't create create exercise: %w", err)
	}

	return nil
}

func (s *Storage) CheckState(username string) (storage.State, error) {
	q := `SELECT state FROM users WHERE username=?`

	var state storage.State
	if err := s.db.QueryRow(q, username).Scan(&state); err != nil {
		return 0, fmt.Errorf("can't check user state: %w", err)
	}

	return state, nil
}

func (s *Storage) ChangeState(username string, state storage.State) error {
	q := `UPDATE users SET state=? WHERE username=?`

	if _, err := s.db.Exec(q, state, username); err != nil {
		return fmt.Errorf("can't update state: %w", err)
	}

	return nil
}
