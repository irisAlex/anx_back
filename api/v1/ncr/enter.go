package ncr

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SupplierApi
	TypeApi
	ProjectApi
	ManageApi
	ComplaintApi
}

var (
	SupplierService  = service.ServiceGroupApp.NcrServiceGroup.SupplierApiService
	TypeService      = service.ServiceGroupApp.NcrServiceGroup.TypeApiService
	ProjectService   = service.ServiceGroupApp.NcrServiceGroup.ProjectService
	ManageService    = service.ServiceGroupApp.NcrServiceGroup.ManageService
	ComplaintService = service.ServiceGroupApp.NcrServiceGroup.ComplaintService
	MessageService   = service.ServiceGroupApp.NcrServiceGroup.MessageService
)
