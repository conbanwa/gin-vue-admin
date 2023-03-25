package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TypeRouter struct {
}

// InitTypeRouter 初始化 Type 路由信息
func (s *TypeRouter) InitTypeRouter(Router *gin.RouterGroup) {
	type_Router := Router.Group("type_").Use(middleware.OperationRecord())
	type_RouterWithoutRecord := Router.Group("type_")
	var type_Api = v1.ApiGroupApp.GreenApiGroup.TypeApi
	{
		type_Router.POST("createType", type_Api.CreateType)   // 新建Type
		type_Router.DELETE("deleteType", type_Api.DeleteType) // 删除Type
		type_Router.DELETE("deleteTypeByIds", type_Api.DeleteTypeByIds) // 批量删除Type
		type_Router.PUT("updateType", type_Api.UpdateType)    // 更新Type
	}
	{
		type_RouterWithoutRecord.GET("findType", type_Api.FindType)        // 根据ID获取Type
		type_RouterWithoutRecord.GET("getTypeList", type_Api.GetTypeList)  // 获取Type列表
	}
}
