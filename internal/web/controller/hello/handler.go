package hello_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"go.uber.org/zap"
)

type Handler interface {
	i()
	View() core.HandlerFunc
}

type handler struct {
	logger *zap.Logger
	cache  cache.Repo
}

func (h *handler) i() {}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger,
		cache,
	}
}
