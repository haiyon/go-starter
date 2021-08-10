package controller

import (
	"go-starter/common/xhttp"
	"go-starter/internal/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello 用户登录
func (ctrl *Controller) Hello(ctx *gin.Context) {
	var (
		err  error
		body dto.SampleBody
	)

	if err = ctx.ShouldBind(&body); err != nil {
		exception := &xhttp.ResponseException{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		xhttp.Fail(ctx, exception)
		return
	}

	result, _ := ctrl.s.Hello(ctx, body)

	xhttp.Success(ctx, result)
}
