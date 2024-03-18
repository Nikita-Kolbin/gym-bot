package storage

type Storage interface {
	CreateUser(username string) error
	CheckState(username string) (State, error)
	ChangeState(username string, state State) error
	UserIsExists(username string) (bool, error)

	CreateGroup(username, groupName string) error
}

type State int

const (
	Standard State = iota
	PickCreate
	CreateGroup
	PickGroupForNewExercise
)
