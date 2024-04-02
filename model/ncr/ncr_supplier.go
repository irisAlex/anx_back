package ncr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Supplier struct {
	global.GVA_MODEL
	Country string `json:"country" gorm:"comment:城市"` // 城市
	Genre   string `json:"genre" gorm:"comment:类别"`   // 类别
	Name    string `json:"name" gorm:"comment:名称"`    // 名称
}

func (Supplier) TableName() string {
	return "ncr_supplier"
}
