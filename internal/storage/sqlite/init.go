package sqlite

import "fmt"

func (s *Storage) Init() error {
	if err := s.initUsers(); err != nil {
		return fmt.Errorf("can't init table user: %w", err)
	}

	if err := s.initGroups(); err != nil {
		return fmt.Errorf("can't init table category: %w", err)
	}

	if err := s.initTrainings(); err != nil {
		return fmt.Errorf("can't init table training: %w", err)
	}

	if err := s.initExercises(); err != nil {
		return fmt.Errorf("can't init table exercise: %w", err)
	}

	if err := s.initSets(); err != nil {
		return fmt.Errorf("can't init table set: %w", err)
	}

	if err := s.initCreateExercises(); err != nil {
		return fmt.Errorf("can't init table create exercise: %w", err)
	}

	return nil
}

func (s *Storage) initUsers() error {
	q := `CREATE TABLE IF NOT EXISTS users (
            username TEXT PRIMARY KEY,
            in_train INTEGER,
            state INTEGER
    	)`

	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initCreateExercises() error {
	q := `CREATE TABLE IF NOT EXISTS create_exercises (
            username TEXT PRIMARY KEY,
            group_id INTEGER,
            exercise_name TEXT,
            FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE
    	)`

	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initGroups() error {
	q := `CREATE TABLE IF NOT EXISTS groups (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			group_name TEXT,
			FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initTrainings() error {
	q := `CREATE TABLE IF NOT EXISTS trainings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT,
			description TEXT,
			start TEXT,
			end TEXT,
			FOREIGN KEY (username) REFERENCES users (username) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initExercises() error {
	q := `CREATE TABLE IF NOT EXISTS exercises (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			group_id  INTEGER,
			name TEXT,
			
			total_weight REAL,
			total_count INTEGER,
			
			weight_record REAL,
			count_record INTEGER,
			tonnage_record REAL,
			
			FOREIGN KEY (group_id) REFERENCES groups (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}

func (s *Storage) initSets() error {
	q := `CREATE TABLE IF NOT EXISTS sets (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			training_id  INTEGER,
			exercise_id  INTEGER,
			weight REAL,
			count INTEGER,
			FOREIGN KEY (training_id) REFERENCES trainings (id) ON DELETE CASCADE
			FOREIGN KEY (exercise_id) REFERENCES exercises (id) ON DELETE CASCADE
    	)`
	_, err := s.db.Exec(q)

	return err
}
