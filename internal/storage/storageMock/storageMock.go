package storageMock

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

func (s *Storage) Foo() {
	_ = 69
}
