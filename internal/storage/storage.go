package storage

type Storage interface {
	CreateUser(username string) error
	CheckState(username string) (State, error)
	ChangeState(username string, state State) error

	CreateGroup(username, groupName string) error
}

type State int

const (
	Standard State = iota
	CreateGroup
)
