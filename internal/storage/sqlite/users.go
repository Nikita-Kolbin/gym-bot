package sqlite

import "fmt"

func (s *Storage) CreateUser(username string) error {
	q := `INSERT INTO users (username, state) VALUES (?, ?)`

	ok, err := s.UserIsExists(username)
	if err != nil {
		return fmt.Errorf("can't create user: %w", err)
	}
	if ok {
		return nil
	}

	if _, err = s.db.Exec(q, username, 0); err != nil {
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
