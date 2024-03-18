package sqlite

import "fmt"

func (s *Storage) CreateExercise(exerciseName string, groupID int) error {
	q := `INSERT INTO exercises (name, group_id) VALUES (?, ?)`

	ok, err := s.ExerciseIsExists(exerciseName, groupID)
	if err != nil {
		return fmt.Errorf("can't create exercise: %w", err)
	}
	if ok {
		return nil
	}

	if _, err = s.db.Exec(q, exerciseName, groupID); err != nil {
		return fmt.Errorf("can't create exercise: %w", err)
	}

	return nil
}

func (s *Storage) createSupportExercise(username string) error {
	q := `INSERT INTO support_exercises (username) VALUES (?)`

	if _, err := s.db.Exec(q, username); err != nil {
		return fmt.Errorf("can't create create exercise: %w", err)
	}

	return nil
}

func (s *Storage) ExerciseIsExists(exerciseName string, groupID int) (bool, error) {
	q := `SELECT COUNT(*) FROM exercises WHERE group_id = ? AND name = ?`

	var count int
	if err := s.db.QueryRow(q, groupID, exerciseName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check exercise exists: %w", err)
	}

	return count > 0, nil
}
