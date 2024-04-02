package ncr

import (
	"fmt"

	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SupplierRouter struct{}

func (s *SupplierRouter) InitSupplierApiRouter(Router *gin.RouterGroup) {
	fmt.Println(54555)
	supplierApiRouter := Router.Group("supplier").Use(middleware.OperationRecord())
	//supplierApiRouterWithoutRecord := Router.Group("supplier")

	supplierRouterApi := v1.ApiGroupApp.NcrApiGroup.SupplierApi
	{
		supplierApiRouter.POST("createSupplierApi", supplierRouterApi.CreateSupplierApi) // 创建Api
		// supplierApiRouter.POST("deleteApi", supplierRouterApi.DeleteApi)               // 删除Api
		// supplierApiRouter.POST("getApiById", supplierRouterApi.GetApiById)             // 获取单条Api消息
		// supplierApiRouter.POST("updateApi", supplierRouterApi.UpdateApi)               // 更新api
		// supplierApiRouter.DELETE("deleteApisByIds", supplierRouterApi.DeleteApisByIds) // 删除选中api
	}
	// {
	// 	supplierApiRouterWithoutRecord.POST("getAllApis", supplierRouterApi.GetAllApis) // 获取所有api
	// 	supplierApiRouterWithoutRecord.POST("getApiList", supplierRouterApi.GetApiList) // 获取Api列表
	// }
}
