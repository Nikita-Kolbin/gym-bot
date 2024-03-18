package sqlite

import "fmt"

func (s *Storage) CreateGroup(username, groupName string) error {
	q := `INSERT INTO groups (username, group_name) VALUES (?, ?)`

	ok, err := s.GroupIsExists(username, groupName)
	if err != nil {
		return fmt.Errorf("can't create group: %w", err)
	}
	if ok {
		return nil
	}

	if _, err = s.db.Exec(q, username, groupName); err != nil {
		return fmt.Errorf("can't create group: %w", err)
	}

	return nil
}

func (s *Storage) GroupIsExists(username, groupName string) (bool, error) {
	q := `SELECT COUNT(*) FROM groups WHERE username = ? AND group_name = ?`

	var count int
	if err := s.db.QueryRow(q, username, groupName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check group exists: %w", err)
	}

	return count > 0, nil
}

func (s *Storage) CheckGroupID(username string, groupName string) (int, error) {
	q := `SELECT id FROM groups WHERE username = ? AND group_name = ?`

	var id int
	if err := s.db.QueryRow(q, username, groupName).Scan(&id); err != nil {
		return 0, fmt.Errorf("can't check group id: %w", err)
	}

	return id, nil
}
