package middleware

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/nedpals/supabase-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth struct {
	supabaseClient *supabase.Client
}

func New(supabaseClient *supabase.Client) *Auth {
	return &Auth{supabaseClient}
}

func (a Auth) Middleware(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	authUser, err := a.supabaseClient.Auth.User(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	ctx = logging.InjectFields(ctx, logging.Fields{"auth.sub", authUser.ID})

	// WARNING: In production define your own type to avoid context collisions.
	return context.WithValue(ctx, "user", authUser), nil
}
