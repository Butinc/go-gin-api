package department_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/service/department_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"

	_ "github.com/spf13/cast"
)

type modifyDepartmentInfoRequest struct {
	Id   string `form:"id"`
	Name string `form:"name"` // 名称
}

type modifyDepartmentInfoResponse struct {
	Name string `json:"name"` // 名称
}

// ModifyDepartmentInfo 修改部门信息
// @Summary 修改部门信息
// @Description 修改部门信息
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "名称"
// @Success 200 {object} modifyDepartmentInfoResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin/modify_password [patch]
func (h *handler) ModifyDepartmentInfo() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyDepartmentInfoRequest)
		res := new(modifyDepartmentInfoResponse)
		if err := c.ShouldBindForm(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		ids, err := h.hashids.HashidsDecode(req.Id)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.HashIdsDecodeError,
				code.Text(code.HashIdsDecodeError)).WithErr(err),
			)
			return
		}

		id := int32(ids[0])

		modifyData := new(department_service.ModifyData)
		modifyData.Name = req.Name

		if err := h.departmentService.ModifyDepartmentInfo(c, id, modifyData); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.DepartmentModifyInfoError,
				code.Text(code.DepartmentModifyInfoError)).WithErr(err),
			)
			return
		}

		res.Name = req.Name
		c.Payload(res)
	}
}
