package storageMock

import "log"

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) CreateUser(username string) error {
	log.Printf("storage mock: create user with username %s", username)
	return nil
}
