package department_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/api/service/department_service"
	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/time_parse"

	"github.com/spf13/cast"
)

type listRequest struct {
	Page     int    `form:"page"`      // 第几页
	PageSize int    `form:"page_size"` // 每页显示条数
	Name     string `form:"name"`      // 名称
}

type listData struct {
	Id          int      `json:"id"`           // ID
	HashID      string   `json:"hashid"`       // hashid
	Name        string   `json:"name"`         // 名称
	CreatedAt   string   `json:"created_at"`   // 创建时间
	CreatedUser string   `json:"created_user"` // 创建人
	UpdatedAt   string   `json:"updated_at"`   // 更新时间
	UpdatedUser string   `json:"updated_user"` // 更新人
	Admins      []string `json:"admins"`       // 管理员
}

type listResponse struct {
	List       []listData `json:"list"`
	Pagination struct {
		Total        int `json:"total"`
		CurrentPage  int `json:"current_page"`
		PerPageCount int `json:"per_page_count"`
	} `json:"pagination"`
}

// List 部门列表
// @Summary 部门列表
// @Description 部门列表
// @Tags API.admin
// @Accept multipart/form-data
// @Produce json
// @Param page query int false "第几页"
// @Param page_size query string false "每页显示条数"
// @Param name query string false "部门"
// @Success 200 {object} listResponse
// @Failure 400 {object} code.Failure
// @Router /api/admin [get]
func (h *handler) List() core.HandlerFunc {
	return func(c core.Context) {
		req := new(listRequest)
		res := new(listResponse)
		if err := c.ShouldBindQuery(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		page := req.Page
		if page == 0 {
			page = 1
		}

		pageSize := req.PageSize
		if pageSize == 0 {
			pageSize = 10
		}

		searchData := new(department_service.SearchData)
		searchData.Page = page
		searchData.PageSize = pageSize
		searchData.Name = req.Name

		resListData, err := h.departmentService.PageList(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithErr(err),
			)
			return
		}

		resCountData, err := h.departmentService.PageListCount(c, searchData)
		if err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.AdminListError,
				code.Text(code.AdminListError)).WithErr(err),
			)
			return
		}
		res.Pagination.Total = cast.ToInt(resCountData)
		res.Pagination.PerPageCount = pageSize
		res.Pagination.CurrentPage = page
		res.List = make([]listData, len(resListData))

		for k, v := range resListData {
			hashId, err := h.hashids.HashidsEncode([]int{cast.ToInt(v.Id)})
			if err != nil {
				c.AbortWithError(errno.NewError(
					http.StatusBadRequest,
					code.HashIdsEncodeError,
					code.Text(code.HashIdsEncodeError)).WithErr(err),
				)
				return
			}

			data := listData{
				Id:          cast.ToInt(v.Id),
				HashID:      hashId,
				Name:        v.Name,
				CreatedAt:   v.CreatedAt.Format(time_parse.CSTLayout),
				CreatedUser: v.CreatedUser,
				UpdatedAt:   v.UpdatedAt.Format(time_parse.CSTLayout),
				UpdatedUser: v.UpdatedUser,
			}

			for _, admin := range v.Admins {
				data.Admins = append(data.Admins, admin.Username)
			}

			res.List[k] = data
		}

		c.Payload(res)
	}
}
