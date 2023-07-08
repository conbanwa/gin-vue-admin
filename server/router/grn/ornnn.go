package grn

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrnRouter struct {
}

// InitOrnRouter 初始化 Orn 路由信息
func (s *OrnRouter) InitOrnRouter(Router *gin.RouterGroup) {
	ornnRouter := Router.Group("ornn").Use(middleware.OperationRecord())
	ornnRouterWithoutRecord := Router.Group("ornn")
	var ornnApi = v1.ApiGroupApp.GrnApiGroup.OrnApi
	{
		ornnRouter.POST("createOrn", ornnApi.CreateOrn)   // 新建Orn
		ornnRouter.DELETE("deleteOrn", ornnApi.DeleteOrn) // 删除Orn
		ornnRouter.DELETE("deleteOrnByIds", ornnApi.DeleteOrnByIds) // 批量删除Orn
		ornnRouter.PUT("updateOrn", ornnApi.UpdateOrn)    // 更新Orn
	}
	{
		ornnRouterWithoutRecord.GET("findOrn", ornnApi.FindOrn)        // 根据ID获取Orn
		ornnRouterWithoutRecord.GET("getOrnList", ornnApi.GetOrnList)  // 获取Orn列表
	}
}
