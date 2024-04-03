package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
)

type TypeResponse struct {
	Ncr ncr.TypeM `json:"typem"`
}

// type SysAPIListResponse struct {
// 	Apis []system.SysApi `json:"apis"`
// }
