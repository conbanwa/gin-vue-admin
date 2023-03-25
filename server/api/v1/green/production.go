package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/green"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    greenReq "github.com/flipped-aurora/gin-vue-admin/server/model/green/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ProductionApi struct {
}

var productionService = service.ServiceGroupApp.GreenServiceGroup.ProductionService


// CreateProduction 创建Production
// @Tags Production
// @Summary 创建Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Production true "创建Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /production/createProduction [post]
func (productionApi *ProductionApi) CreateProduction(c *gin.Context) {
	var production green.Production
	err := c.ShouldBindJSON(&production)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    production.CreatedBy = utils.GetUserID(c)
	if err := productionService.CreateProduction(production); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProduction 删除Production
// @Tags Production
// @Summary 删除Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Production true "删除Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /production/deleteProduction [delete]
func (productionApi *ProductionApi) DeleteProduction(c *gin.Context) {
	var production green.Production
	err := c.ShouldBindJSON(&production)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    production.DeletedBy = utils.GetUserID(c)
	if err := productionService.DeleteProduction(production); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProductionByIds 批量删除Production
// @Tags Production
// @Summary 批量删除Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /production/deleteProductionByIds [delete]
func (productionApi *ProductionApi) DeleteProductionByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := productionService.DeleteProductionByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProduction 更新Production
// @Tags Production
// @Summary 更新Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Production true "更新Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /production/updateProduction [put]
func (productionApi *ProductionApi) UpdateProduction(c *gin.Context) {
	var production green.Production
	err := c.ShouldBindJSON(&production)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    production.UpdatedBy = utils.GetUserID(c)
	if err := productionService.UpdateProduction(production); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProduction 用id查询Production
// @Tags Production
// @Summary 用id查询Production
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query green.Production true "用id查询Production"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /production/findProduction [get]
func (productionApi *ProductionApi) FindProduction(c *gin.Context) {
	var production green.Production
	err := c.ShouldBindQuery(&production)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reproduction, err := productionService.GetProduction(production.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproduction": reproduction}, c)
	}
}

// GetProductionList 分页获取Production列表
// @Tags Production
// @Summary 分页获取Production列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query greenReq.ProductionSearch true "分页获取Production列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /production/getProductionList [get]
func (productionApi *ProductionApi) GetProductionList(c *gin.Context) {
	var pageInfo greenReq.ProductionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := productionService.GetProductionInfoList(pageInfo); err != nil {
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
