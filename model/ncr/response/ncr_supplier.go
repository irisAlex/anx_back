package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
)

type SupplierResponse struct {
	Ncr ncr.Supplier `json:"supplier"`
}

// type SysAPIListResponse struct {
// 	Apis []system.SysApi `json:"apis"`
// }
