package ncr

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

// 供应商 router
type SupplierRouter struct{}

func (s *SupplierRouter) InitSupplierRouter(Router *gin.RouterGroup) {
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

// 类型 router
type TypeRouter struct{}

func (s *TypeRouter) InitTypeRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("type").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("type")

	typeRouterApi := v1.ApiGroupApp.NcrApiGroup.TypeApi
	{
		typeApiRouter.POST("createType", typeRouterApi.CreateTypeApi) // 创建Api
		typeApiRouter.POST("deleteType", typeRouterApi.DeleteType)    // 删除Api
		typeApiRouter.POST("getTypeById", typeRouterApi.GetTypeById)  // 获取单条Api消息
		typeApiRouter.POST("updateType", typeRouterApi.UpdateType)    // 更新api
		// supplierApiRouter.DELETE("deleteApisByIds", supplierRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllTypeList", typeRouterApi.GetTypeApiList) // 获取所有api
		//supplierApiRouterWithoutRecord.POST("getSupplierList", supplierRouterApi.GetSupplierList) // 获取Api列表
	}
}

//项目

type ProjectRouter struct{}

func (s *TypeRouter) InitProjectRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("project").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("project")

	typeRouterApi := v1.ApiGroupApp.NcrApiGroup.ProjectApi
	{
		typeApiRouter.POST("createProject", typeRouterApi.CreateProject)   // 创建Api
		typeApiRouter.POST("deleteProject", typeRouterApi.DeleteProject)   // 删除Api
		typeApiRouter.POST("getProjectById", typeRouterApi.GetProjectById) // 获取单条Api消息
		typeApiRouter.POST("updateProject", typeRouterApi.UpdateProject)   // 更新api
		// supplierApiRouter.DELETE("deleteApisByIds", supplierRouterApi.DeleteApisByIds) // 删除选中api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllProjectList", typeRouterApi.GetProjectList) // 获取所有api
	}
}

//manage
//项目

type ManageRouter struct{}

func (s *ManageRouter) InitManageRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("manage").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("manage")

	typeRouterApi := v1.ApiGroupApp.NcrApiGroup.ManageApi
	{
		typeApiRouter.POST("createManage", typeRouterApi.CreateManage)   // 创建Api
		typeApiRouter.POST("deleteManage", typeRouterApi.DeleteManage)   // 删除Api
		typeApiRouter.POST("getManageById", typeRouterApi.GetManageById) // 获取单条Api消息
		typeApiRouter.POST("updateManage", typeRouterApi.UpdateManage)   // 更新api
		typeApiRouter.POST("updateSetStatus", typeRouterApi.SetStatus)   // 更新api
		typeApiRouter.POST("updateParts", typeRouterApi.UpdateParts)     // 更新api
		typeApiRouter.POST("updatePassDate", typeRouterApi.SetPassDate)  // 更新放行时间
	}
	{
		typeApiRouterWithoutRecord.POST("getAllManageList", typeRouterApi.GetManageList) // 获取所有api
	}

}

type ComplaintRouter struct{}

func (s *ComplaintRouter) InitComplaintRouter(Router *gin.RouterGroup) {
	typeApiRouter := Router.Group("complaint").Use(middleware.OperationRecord())
	typeApiRouterWithoutRecord := Router.Group("complaint")

	typeRouterApi := v1.ApiGroupApp.NcrApiGroup.ComplaintApi
	{
		typeApiRouter.POST("createComplaint", typeRouterApi.CreateComplaint)   // 创建Api
		typeApiRouter.POST("deleteComplaint", typeRouterApi.DeleteComplaint)   // 删除Api
		typeApiRouter.POST("getComplaintById", typeRouterApi.GetComplaintById) // 获取单条Api消息
		typeApiRouter.POST("updateComplaint", typeRouterApi.UpdateComplaint)   // 更新api
		typeApiRouter.POST("updateSetStatus", typeRouterApi.SetStatus)         // 更新api
	}
	{
		typeApiRouterWithoutRecord.POST("getAllComplaintList", typeRouterApi.GetComplaintList) // 获取所有api
	}

}
