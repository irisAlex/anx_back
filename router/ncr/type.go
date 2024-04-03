package ncr

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TypeRouter struct{}

func (s *TypeRouter) InitTypeApiRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("type").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("type")

	typeRouterApi := v1.ApiGroupApp.NcrApiGroup.SupplierApi
	{
		typeApiRouter.POST("createSupplierApi", typeRouterApi.CreateSupplierApi) // 创建Api
		typeApiRouter.POST("deleteSupplier", typeRouterApi.DeleteSupplier)       // 删除Api
		typeApiRouter.POST("getSupplierById", typeRouterApi.GetSupplierById)     // 获取单条Api消息
		typeApiRouter.POST("updateSupplier", typeRouterApi.UpdateSupplier)       // 更新api
		// supplierApiRouter.DELETE("deleteApisByIds", supplierRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllSupplierList", typeRouterApi.GetSupplierApiList) // 获取所有api
		//supplierApiRouterWithoutRecord.POST("getSupplierList", supplierRouterApi.GetSupplierList) // 获取Api列表
	}
}
