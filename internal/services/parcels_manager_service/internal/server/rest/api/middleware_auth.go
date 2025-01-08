package api

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subliker/track-parcel-service/internal/pkg/client/grpc/account/manager"
	"github.com/subliker/track-parcel-service/internal/pkg/domain/model"
	"github.com/subliker/track-parcel-service/internal/pkg/gen/account/managerpb"
)

type contextKey uint

const (
	contextKeyManagerTID contextKey = iota
)

func (s *Server) authApiTokenMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// get api token
			apiToken := r.Header.Get("Authorization")
			if apiToken == "" {
				http.Error(w, "api token is not set", http.StatusUnauthorized)
				return
			}

			// auth by api token
			res, err := s.managerClient.AuthApiToken(context.Background(), &managerpb.AuthApiTokenRequest{
				ManagerApiToken: apiToken,
			})
			if err == manager.ErrManagerNotFound {
				http.Error(w, "api token is not valid", http.StatusBadRequest)
				return
			}
			if err != nil {
				http.Error(w, "internal error", http.StatusInternalServerError)
				return
			}

			// add authorized manager id in context
			ctx := context.WithValue(r.Context(), contextKeyManagerTID, model.TelegramID(res.ManagerTelegramId))
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
