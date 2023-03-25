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

type GardenApi struct {
}

var gardenService = service.ServiceGroupApp.GreenServiceGroup.GardenService


// CreateGarden 创建Garden
// @Tags Garden
// @Summary 创建Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Garden true "创建Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garden/createGarden [post]
func (gardenApi *GardenApi) CreateGarden(c *gin.Context) {
	var garden green.Garden
	err := c.ShouldBindJSON(&garden)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    garden.CreatedBy = utils.GetUserID(c)
	if err := gardenService.CreateGarden(garden); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGarden 删除Garden
// @Tags Garden
// @Summary 删除Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Garden true "删除Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /garden/deleteGarden [delete]
func (gardenApi *GardenApi) DeleteGarden(c *gin.Context) {
	var garden green.Garden
	err := c.ShouldBindJSON(&garden)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    garden.DeletedBy = utils.GetUserID(c)
	if err := gardenService.DeleteGarden(garden); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGardenByIds 批量删除Garden
// @Tags Garden
// @Summary 批量删除Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /garden/deleteGardenByIds [delete]
func (gardenApi *GardenApi) DeleteGardenByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := gardenService.DeleteGardenByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGarden 更新Garden
// @Tags Garden
// @Summary 更新Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Garden true "更新Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /garden/updateGarden [put]
func (gardenApi *GardenApi) UpdateGarden(c *gin.Context) {
	var garden green.Garden
	err := c.ShouldBindJSON(&garden)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    garden.UpdatedBy = utils.GetUserID(c)
	if err := gardenService.UpdateGarden(garden); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGarden 用id查询Garden
// @Tags Garden
// @Summary 用id查询Garden
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query green.Garden true "用id查询Garden"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /garden/findGarden [get]
func (gardenApi *GardenApi) FindGarden(c *gin.Context) {
	var garden green.Garden
	err := c.ShouldBindQuery(&garden)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if regarden, err := gardenService.GetGarden(garden.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regarden": regarden}, c)
	}
}

// GetGardenList 分页获取Garden列表
// @Tags Garden
// @Summary 分页获取Garden列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query greenReq.GardenSearch true "分页获取Garden列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /garden/getGardenList [get]
func (gardenApi *GardenApi) GetGardenList(c *gin.Context) {
	var pageInfo greenReq.GardenSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gardenService.GetGardenInfoList(pageInfo); err != nil {
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
