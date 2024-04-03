package ncr

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SupplierApi
	TypeApi
}

var (
	SupplierService = service.ServiceGroupApp.NcrServiceGroup.SupplierApiService
	TypeService     = service.ServiceGroupApp.NcrServiceGroup.SupplierApiService
)
