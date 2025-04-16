package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-basic/internal/service"
	"github.com/go-nunu/nunu-layout-basic/pkg/helper/result"
	"go.uber.org/zap"
)

// @wire:Handler
func NewUserHandler(handler *Handler,
	userService service.UserService,
) *UserHandler {
	return &UserHandler{
		Handler:     handler,
		userService: userService,
	}
}

type UserHandler struct {
	*Handler
	userService service.UserService
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	var params struct {
		Id int64 `form:"id" binding:"required"`
	}
	if err := ctx.ShouldBind(&params); err != nil {
		result.FailByErr(ctx, err)
		return
	}

	user, err := h.userService.GetUserById(params.Id)
	h.logger.Info("GetUserByID", zap.Any("user", user))
	if err != nil {
		result.FailByErr(ctx, err)
		return
	}
	result.Success(ctx, user)
	//resp.HandleSuccess(ctx, user)
}

func (h *UserHandler) UpdateUser(ctx *gin.Context) {
	result.Success(ctx, "")
}
