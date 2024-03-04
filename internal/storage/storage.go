package storage

type Storage interface {
	CreateUser(username string) error
}
