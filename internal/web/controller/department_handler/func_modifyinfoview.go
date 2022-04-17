package department_handler

import (
	"net/http"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
)

type modifyInfoRequest struct {
	Id string `uri:"id"`
}

type modifyInfoResponse struct {
	HashId string `json:"hash_id"`
}

func (h *handler) ModifyInfoView() core.HandlerFunc {
	return func(c core.Context) {
		req := new(modifyInfoRequest)
		res := new(modifyInfoResponse)

		if err := c.ShouldBindURI(req); err != nil {
			c.AbortWithError(errno.NewError(
				http.StatusBadRequest,
				code.ParamBindError,
				code.Text(code.ParamBindError)).WithErr(err),
			)
			return
		}

		res.HashId = req.Id
		c.HTML("department_modifyinfo", res)
	}
}
