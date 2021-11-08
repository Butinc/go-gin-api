package department_handler

import (
	_ "encoding/json"
	"net/http"

	_ "github.com/xinliangnote/go-gin-api/configs"
	"github.com/xinliangnote/go-gin-api/internal/api/service/department_service"
	_ "github.com/xinliangnote/go-gin-api/internal/pkg/cache"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	_ "github.com/xinliangnote/go-gin-api/internal/pkg/password"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	_ "github.com/spf13/cast"
)

type detailRequest struct {
	Id int32 `uri:"id"`
}
type detailResponse struct {
	Name string `json:"name"` // 用户名
}

// Detail 部门详情
// @Summary 部门详情
// @Description 部门详情
// @Tags API.department
// @Accept json
// @Produce json
// @Success 200 {object} detailResponse
// @Failure 400 {object} code.Failure
// @Router /api/department/info [get]
func (h *handler) Detail() core.HandlerFunc {
	return func(c core.Context) {
		req := new(detailRequest)
		res := new(detailResponse)

		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}
		searchOneData := new(department_service.SearchOneData)
		searchOneData.Id = req.Id

		info, err := h.departmentService.Detail(c, searchOneData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminDetailError,
				code.Text(code.AdminDetailError)).WithErr(err),
			)
			return
		}

		res.Name = info.Name
		c.Payload(res)
	}
}
