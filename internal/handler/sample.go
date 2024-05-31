package handler

import (
	"go-starter/internal/data/structs"
	"go-starter/pkg/resp"
	"go-starter/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Hello is a handler for the "Hello" endpoint.
// @Produce json
// @Success 200 {object} resp.Exception "ok"
// @Failure 400 {object} resp.Exception "bad request"
// @Router /sample/hello [get]
func (h *Handler) Hello(ctx *gin.Context) {
	var (
		err  error
		body structs.Sample
	)

	if err = ctx.ShouldBind(&body); err != nil {
		exception := &resp.Exception{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}
		resp.Fail(ctx.Writer, exception)
		return
	}

	result, err := h.svc.Hello(ctx, body)
	if validator.IsNotNil(err) {
		resp.Fail(ctx.Writer, result)
		return
	}

	resp.Success(ctx.Writer, result)
}
