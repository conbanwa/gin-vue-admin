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

type TypeApi struct {
}

var type_Service = service.ServiceGroupApp.GreenServiceGroup.TypeService


// CreateType 创建Type
// @Tags Type
// @Summary 创建Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Type true "创建Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /type_/createType [post]
func (type_Api *TypeApi) CreateType(c *gin.Context) {
	var type_ green.Type
	err := c.ShouldBindJSON(&type_)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    type_.CreatedBy = utils.GetUserID(c)
	if err := type_Service.CreateType(type_); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteType 删除Type
// @Tags Type
// @Summary 删除Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Type true "删除Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /type_/deleteType [delete]
func (type_Api *TypeApi) DeleteType(c *gin.Context) {
	var type_ green.Type
	err := c.ShouldBindJSON(&type_)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    type_.DeletedBy = utils.GetUserID(c)
	if err := type_Service.DeleteType(type_); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTypeByIds 批量删除Type
// @Tags Type
// @Summary 批量删除Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /type_/deleteTypeByIds [delete]
func (type_Api *TypeApi) DeleteTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    deletedBy := utils.GetUserID(c)
	if err := type_Service.DeleteTypeByIds(IDS,deletedBy); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateType 更新Type
// @Tags Type
// @Summary 更新Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body green.Type true "更新Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /type_/updateType [put]
func (type_Api *TypeApi) UpdateType(c *gin.Context) {
	var type_ green.Type
	err := c.ShouldBindJSON(&type_)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    type_.UpdatedBy = utils.GetUserID(c)
	if err := type_Service.UpdateType(type_); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindType 用id查询Type
// @Tags Type
// @Summary 用id查询Type
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query green.Type true "用id查询Type"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /type_/findType [get]
func (type_Api *TypeApi) FindType(c *gin.Context) {
	var type_ green.Type
	err := c.ShouldBindQuery(&type_)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retype_, err := type_Service.GetType(type_.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retype_": retype_}, c)
	}
}

// GetTypeList 分页获取Type列表
// @Tags Type
// @Summary 分页获取Type列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query greenReq.TypeSearch true "分页获取Type列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /type_/getTypeList [get]
func (type_Api *TypeApi) GetTypeList(c *gin.Context) {
	var pageInfo greenReq.TypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := type_Service.GetTypeInfoList(pageInfo); err != nil {
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
