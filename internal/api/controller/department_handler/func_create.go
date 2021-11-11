package department_handler

import (
	"github.com/xinliangnote/go-gin-api/internal/api/service/department_service"
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/validation"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type createRequest struct {
	Name string `form:"name" binding:"required"` // 用户名
}

type createResponse struct {
	Id int32 `json:"id"` // 主键ID
}

// Create 新增部门
// @Summary 新增部门
// @Description 新增部门
// @Tags API.department
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "部门"
// @Success 200 {object} createResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [post]
func (h *handler) Create() core.HandlerFunc {
	return func(c core.Context) {
		req := new(createRequest)
		res := new(createResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithErr(err),
			)
			return
		}

		createData := new(department_service.CreateDepartmentData)
		createData.Name = req.Name

		id, err := h.departmentService.Create(c, createData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.DepartmentCreateError,
				code.Text(code.DepartmentCreateError)).WithErr(err),
			)
			return
		}

		res.Id = id
		c.Payload(res)
	}
}
