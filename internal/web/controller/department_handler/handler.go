package department_handler

import (
	"github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/department_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/db"
	"github.com/xinliangnote/go-gin-api/pkg/hash"

	"go.uber.org/zap"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	AddView() core.HandlerFunc
	ListView() core.HandlerFunc
	ModifyInfoView() core.HandlerFunc
}

type handler struct {
	db                db.Repo
	logger            *zap.Logger
	cache             cache.Repo
	hashids           hash.Hash
	departmentService department_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		db:                db,
		hashids:           hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		departmentService: department_service.New(db, cache),
	}
}

func (h *handler) i() {}
