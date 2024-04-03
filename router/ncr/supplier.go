package ncr

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplierRouter struct{}

func (s *SupplierRouter) InitSupplierApiRouter(Router *gin.RouterGroup) {
	supplierApiRouter := Router.Group("supplier").Use(middleware.OperationRecord())
	supplierApiRouterWithoutRecord := Router.Group("supplier")

	supplierRouterApi := v1.ApiGroupApp.NcrApiGroup.SupplierApi
	{
		supplierApiRouter.POST("createSupplierApi", supplierRouterApi.CreateSupplierApi) // 创建Api
		supplierApiRouter.POST("deleteSupplier", supplierRouterApi.DeleteSupplier)       // 删除Api
		supplierApiRouter.POST("getSupplierById", supplierRouterApi.GetSupplierById)     // 获取单条Api消息
		supplierApiRouter.POST("updateSupplier", supplierRouterApi.UpdateSupplier)       // 更新api
		// supplierApiRouter.DELETE("deleteApisByIds", supplierRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		supplierApiRouterWithoutRecord.POST("getAllSupplierList", supplierRouterApi.GetSupplierApiList) // 获取所有api
		//supplierApiRouterWithoutRecord.POST("getSupplierList", supplierRouterApi.GetSupplierList) // 获取Api列表
	}
}
