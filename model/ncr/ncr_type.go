package ncr

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type TypeM struct {
	global.GVA_MODEL
	Genre string `json:"genre" gorm:"comment:类别"` // 类别
	Name  string `json:"name" gorm:"comment:名称"`  // 名称
}

func (TypeM) TableName() string {
	return "ncr_type"
}
