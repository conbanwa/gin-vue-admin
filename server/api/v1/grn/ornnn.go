package grn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/grn"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    grnReq "github.com/flipped-aurora/gin-vue-admin/server/model/grn/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type OrnApi struct {
}

var ornnService = service.ServiceGroupApp.GrnServiceGroup.OrnService


// CreateOrn 创建Orn
// @Tags Orn
// @Summary 创建Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body grn.Orn true "创建Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ornn/createOrn [post]
func (ornnApi *OrnApi) CreateOrn(c *gin.Context) {
	var ornn grn.Orn
	err := c.ShouldBindJSON(&ornn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ornnService.CreateOrn(&ornn); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrn 删除Orn
// @Tags Orn
// @Summary 删除Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body grn.Orn true "删除Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ornn/deleteOrn [delete]
func (ornnApi *OrnApi) DeleteOrn(c *gin.Context) {
	var ornn grn.Orn
	err := c.ShouldBindJSON(&ornn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ornnService.DeleteOrn(ornn); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrnByIds 批量删除Orn
// @Tags Orn
// @Summary 批量删除Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ornn/deleteOrnByIds [delete]
func (ornnApi *OrnApi) DeleteOrnByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ornnService.DeleteOrnByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrn 更新Orn
// @Tags Orn
// @Summary 更新Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body grn.Orn true "更新Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ornn/updateOrn [put]
func (ornnApi *OrnApi) UpdateOrn(c *gin.Context) {
	var ornn grn.Orn
	err := c.ShouldBindJSON(&ornn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ornnService.UpdateOrn(ornn); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrn 用id查询Orn
// @Tags Orn
// @Summary 用id查询Orn
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query grn.Orn true "用id查询Orn"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ornn/findOrn [get]
func (ornnApi *OrnApi) FindOrn(c *gin.Context) {
	var ornn grn.Orn
	err := c.ShouldBindQuery(&ornn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reornn, err := ornnService.GetOrn(ornn.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reornn": reornn}, c)
	}
}

// GetOrnList 分页获取Orn列表
// @Tags Orn
// @Summary 分页获取Orn列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query grnReq.OrnSearch true "分页获取Orn列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ornn/getOrnList [get]
func (ornnApi *OrnApi) GetOrnList(c *gin.Context) {
	var pageInfo grnReq.OrnSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ornnService.GetOrnInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
