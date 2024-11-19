package store

type Store interface {
	Close() error
	User() UserRepository
	Manager() ManagerRepository
}
