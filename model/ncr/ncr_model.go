package ncr

import (
	"time"

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

type TypeM struct {
	global.GVA_MODEL
	Genre string `json:"genre" gorm:"comment:类别"` // 类别
	Name  string `json:"name" gorm:"comment:名称"`  // 名称
}

func (TypeM) TableName() string {
	return "ncr_type"
}

type Project struct {
	global.GVA_MODEL
	Name      string `json:"name" gorm:"comment:类别"`      // 类别
	Period    int64  `json:"period" gorm:"comment:名称"`    // 名称
	Principal string `json:"principal" gorm:"comment:名称"` // 名称
	Describe  string `json:"describe" gorm:"comment:名称"`  // 名称
	Priority  string `json:"priority" gorm:"comment:名称"`  // 名称
}

func (Project) TableName() string {
	return "ncr_project"
}

type Manage struct {
	global.GVA_MODEL
	Serialnumber                  string    `json:"serialnumber" gorm:"comment:类别"`                  // 类别
	Department                    string    `json:"department" gorm:"comment:类别"`                    // 类别
	Type_M                        string    `json:"type_m" gorm:"comment:类别"`                        // 类别
	Category                      string    `json:"category" gorm:"comment:类别"`                      // 类别
	Project                       string    `json:"project" gorm:"comment:类别"`                       // 类别
	Checkout_Name                 string    `json:"checkout_name" gorm:"comment:类别"`                 // 类别
	Checkout_Number               int64     `json:"checkout_number" gorm:"comment:名称"`               // 名称
	Graph_Number                  int64     `json:"graph_number" gorm:"comment:名称"`                  // 名称
	Version_Number                int64     `json:"version_number" gorm:"comment:名称"`                // 名称
	Purchase_Order                int64     `json:"purchase_order" gorm:"comment:名称"`                // 名称
	Production_Order              int64     `json:"production_order" gorm:"comment:名称"`              // 名称
	Delivery_Order                int64     `json:"delivery_order" gorm:"comment:名称"`                // 名称
	Packages_Number               int64     `json:"packages_number" gorm:"comment:名称"`               // 名称
	Reject_Packages_Number        int64     `json:"reject_packages_number" gorm:"comment:名称"`        // 名称
	Sample_Checkout_Number        int64     `json:"sample_checkout_number" gorm:"comment:名称"`        // 名称
	Reject_Sample_Checkout_Number int64     `json:"reject_sample_checkout_number" gorm:"comment:名称"` // 名称
	Supplier                      string    `json:"supplier" gorm:"comment:类别"`                      // 类别
	Checkout_Date                 time.Time `json:"checkout_date" gorm:"comment:名称"`                 // 名称
	Describe                      string    `json:"describe" gorm:"comment:类别"`                      // 类别
	Photograph                    string    `json:"photograph" gorm:"comment:类别"`                    // 类别
	Process_Mode                  string    `json:"process_mode" gorm:"comment:类别"`                  // 类别
	Duty_Department               string    `json:"duty_department" gorm:"comment:类别"`               // 类别
	Cause_Desc                    string    `json:"cause_desc" gorm:"comment:类别"`                    // 类别
	Disposal_Concept              string    `json:"disposal_concept" gorm:"comment:类别"`              // 类别
	Rework_Number                 int64     `json:"rework_number" gorm:"comment:名称"`                 // 名称
	Rework_Man_Hour               int64     `json:"rework_man_hour" gorm:"comment:名称"`               // 名称
	Rework_Quantities             string    `json:"rework_quantities" gorm:"comment:类别"`             // 类别
	Rework_Process                string    `json:"rework_process" gorm:"comment:类别"`                // 类别
	Rework_Plan_Date              time.Time `json:"rework_plan_date" gorm:"comment:名称"`              // 名称,
	Rework_Desc                   string    `json:"rework_desc" gorm:"comment:类别"`                   // 类别
	Rework_Attachment             string    `json:"rework_attachment" gorm:"comment:类别"`             // 类别
	Operation_Type                int64     `json:"operation_type" gorm:"comment:名称"`                // 名称
}

func (Manage) TableName() string {
	return "ncr_product_manage"
}
