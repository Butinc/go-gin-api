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

	// Create 新增部门
	// @Tags API.department
	// @Router /api/department [post]
	Create() core.HandlerFunc

	// List 部门列表
	// @Tags API.department
	// @Router /api/department [get]
	List() core.HandlerFunc

	// Detail 部门信息
	// @Tags API.department
	// @Router /api/department/info [get]
	Detail() core.HandlerFunc

	// ModifyDepartmentInfo 修改部门信息
	// @Tags API.department
	// @Router /api/department/offline [patch]
	ModifyDepartmentInfo() core.HandlerFunc

	// Delete 删除管理员
	// @Tags API.admin
	// @Router /api/admin/{id} [delete]
	Delete() core.HandlerFunc
}

type handler struct {
	logger            *zap.Logger
	cache             cache.Repo
	hashids           hash.Hash
	departmentService department_service.Service
}

func New(logger *zap.Logger, db db.Repo, cache cache.Repo) Handler {
	return &handler{
		logger:            logger,
		cache:             cache,
		hashids:           hash.New(configs.Get().HashIds.Secret, configs.Get().HashIds.Length),
		departmentService: department_service.New(db, cache),
	}
}

func (h *handler) i() {}
