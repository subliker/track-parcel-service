package store

// Store is interface for accessing user and manager repositories
//
//go:generate mockgen -source=store.go -destination=mock/store.go -package=mock
type Store interface {
	// User returns user repository
	User() UserRepository
	// Manager returns manager repository
	Manager() ManagerRepository

	// Close closes store
	Close() error
}
