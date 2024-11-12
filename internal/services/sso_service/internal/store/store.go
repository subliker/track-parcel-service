package store

import (
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore/manager"
	"github.com/subliker/track-parcel-service/internal/services/sso_service/internal/store/pgstore/user"
)

type Store interface {
	Close() error
	User() (*user.Repository, error)
	Manager() (*manager.Repository, error)
}
