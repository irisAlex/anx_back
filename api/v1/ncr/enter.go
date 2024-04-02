package ncr

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	SupplierApi
}

var (
	SupplierService = service.ServiceGroupApp.NcrServiceGroup.SupplierApiService
)
