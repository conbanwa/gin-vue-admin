package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GardenRouter struct {
}

// InitGardenRouter 初始化 Garden 路由信息
func (s *GardenRouter) InitGardenRouter(Router *gin.RouterGroup) {
	gardenRouter := Router.Group("garden").Use(middleware.OperationRecord())
	gardenRouterWithoutRecord := Router.Group("garden")
	var gardenApi = v1.ApiGroupApp.GreenApiGroup.GardenApi
	{
		gardenRouter.POST("createGarden", gardenApi.CreateGarden)   // 新建Garden
		gardenRouter.DELETE("deleteGarden", gardenApi.DeleteGarden) // 删除Garden
		gardenRouter.DELETE("deleteGardenByIds", gardenApi.DeleteGardenByIds) // 批量删除Garden
		gardenRouter.PUT("updateGarden", gardenApi.UpdateGarden)    // 更新Garden
	}
	{
		gardenRouterWithoutRecord.GET("findGarden", gardenApi.FindGarden)        // 根据ID获取Garden
		gardenRouterWithoutRecord.GET("getGardenList", gardenApi.GetGardenList)  // 获取Garden列表
	}
}
