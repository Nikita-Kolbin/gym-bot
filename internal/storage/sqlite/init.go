package sqlite

import "fmt"

func (s *Storage) Init() error {
	if err := s.initUser(); err != nil {
		return fmt.Errorf("can't init table user: %w", err)
	}

	if err := s.initCategory(); err != nil {
		return fmt.Errorf("can't init table category: %w", err)
	}

	if err := s.initTraining(); err != nil {
		return fmt.Errorf("can't init table training: %w", err)
	}

	if err := s.initExercise(); err != nil {
		return fmt.Errorf("can't init table exercise: %w", err)
	}

	if err := s.initSet(); err != nil {
		return fmt.Errorf("can't init table set: %w", err)
	}

	return nil
}

func (s *Storage) initUser() error {
	q := `CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT, 
            username TEXT,
            state INTEGER
    	)`

	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initCategory() error {
	q := `CREATE TABLE IF NOT EXISTS categores (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			category_name TEXT,
			user_id  INTEGER,
			FOREIGN KEY (user_id)  REFERENCES user (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initTraining() error {
	q := `CREATE TABLE IF NOT EXISTS trainings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			description TEXT,
			start TEXT,
			end TEXT,
			user_id  INTEGER,
			FOREIGN KEY (user_id)  REFERENCES user (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initExercise() error {
	q := `CREATE TABLE IF NOT EXISTS exercises (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			category_id  INTEGER,
			FOREIGN KEY (category_id)  REFERENCES category (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initSet() error {
	q := `CREATE TABLE IF NOT EXISTS sets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			weight REAL,
			count INTEGER,
			training_id  INTEGER,
			exercise_id  INTEGER,
			FOREIGN KEY (training_id)  REFERENCES training (id) ON DELETE CASCADE
			FOREIGN KEY (exercise_id)  REFERENCES exercse (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}
