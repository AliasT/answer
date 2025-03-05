package controller

import (
	"github.com/apache/answer/internal/base/handler"
	"github.com/apache/answer/internal/base/middleware"
	"github.com/apache/answer/internal/schema"
	tipcommon "github.com/apache/answer/internal/service/tip_common"
	"github.com/gin-gonic/gin"
)

type TipController struct {
	tipCommonService *tipcommon.TipCommonService
}

// NewTipController new tip controller
func NewTipController(tipCommonService *tipcommon.TipCommonService) *TipController {
	return &TipController{
		tipCommonService: tipCommonService,
	}
}

// 给题目打赏
func (tc *TipController) AddTip(ctx *gin.Context) {
	req := &schema.TipReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	req.UserID = middleware.GetLoginUserIDFromContext(ctx)
	req.Amount = 2

	resp, err := tc.tipCommonService.AddTip(ctx, req)

	handler.HandleResponse(ctx, err, resp)

}
