package env

import (
	"context"
	"database/sql"

	"github.com/projects/sys-des/txn-routine/pkg/logger"
)

type envContextKey string

const (
	// EnvCtxKey is the key to set and retrieve Env in context
	EnvCtxKey envContextKey = "env"
)

type Env struct {
	DB *sql.DB
}

func New(ctx context.Context) *Env {
	return &Env{}
}

func setInContext(ctx context.Context, env *Env) context.Context {
	return context.WithValue(ctx, EnvCtxKey, env)
}

func (e *Env) Setup(ctx context.Context) context.Context {
	d := NewDB()
	e.DB = d

	return setInContext(ctx, e)
}

func FromContext(ctx context.Context) *Env {
	env, ok := ctx.Value(EnvCtxKey).(*Env)
	if !ok {
		logger.GetLogger().ErrorContext(ctx, "Failed to get environment from context")
		return nil
	}
	return env
}
