package department_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"

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
	logger *zap.Logger
}

func New(logger *zap.Logger) Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) i() {}
