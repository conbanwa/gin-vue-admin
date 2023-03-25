package green

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductionRouter struct {
}

// InitProductionRouter 初始化 Production 路由信息
func (s *ProductionRouter) InitProductionRouter(Router *gin.RouterGroup) {
	productionRouter := Router.Group("production").Use(middleware.OperationRecord())
	productionRouterWithoutRecord := Router.Group("production")
	var productionApi = v1.ApiGroupApp.GreenApiGroup.ProductionApi
	{
		productionRouter.POST("createProduction", productionApi.CreateProduction)   // 新建Production
		productionRouter.DELETE("deleteProduction", productionApi.DeleteProduction) // 删除Production
		productionRouter.DELETE("deleteProductionByIds", productionApi.DeleteProductionByIds) // 批量删除Production
		productionRouter.PUT("updateProduction", productionApi.UpdateProduction)    // 更新Production
	}
	{
		productionRouterWithoutRecord.GET("findProduction", productionApi.FindProduction)        // 根据ID获取Production
		productionRouterWithoutRecord.GET("getProductionList", productionApi.GetProductionList)  // 获取Production列表
	}
}
