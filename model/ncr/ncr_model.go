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
	Mold                          string    `json:"mold" gorm:"comment:类别"`                          // 类别
	Category                      string    `json:"category" gorm:"comment:类别"`                      // 类别
	Project                       string    `json:"project" gorm:"comment:类别"`                       // 类别
	Checkout_Name                 string    `json:"checkout_name" gorm:"comment:类别"`                 // 类别
	Checkout_Number               string    `json:"checkout_number" gorm:"comment:名称"`               // 名称
	Graph_Number                  string    `json:"graph_number" gorm:"comment:名称"`                  // 名称
	Version_Number                string    `json:"version_number" gorm:"comment:名称"`                // 名称
	Purchase_Order                string    `json:"purchase_order" gorm:"comment:名称"`                // 名称
	Production_Order              string    `json:"production_order" gorm:"comment:名称"`              // 名称
	Delivery_Order                string    `json:"delivery_order" gorm:"comment:名称"`                // 名称
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
	Rework_Attachment             string    `json:"rework_attachment" gorm:"comment:类别"`
	Repair_Plan_Date              string    `json:"repair_plan_date" gorm:"comment:类别"`
	Area                          string    `json:"area" gorm:"comment:类别"`
	Find_Addr                     string    `json:"find_addr" gorm:"comment:类别"`
	Find_Process                  string    `json:"find_process" gorm:"comment:类别"`
	Defect_Problem                string    `json:"defect_problem" gorm:"comment:类别"`
	Station                       string    `json:"station" gorm:"comment:类别"`        //1 返工 2 返修  3 配做 4 让步放行
	Operation_Type                string    `json:"operation_type" gorm:"comment:类别"` //1 返工 2 返修  3 配做 4 让步放行
	Deferred_Date                 time.Time `json:"deferred_date" gorm:"comment:类别"`
	Status                        string    `json:"status" gorm:"comment:类别"`
	Repair_Desc                   string    `json:"repair_desc" gorm:"comment:类别"`
	Repair_Attachment             string    `json:"repair_attachment" gorm:"comment:类别"`
	Parts_Desc                    string    `json:"parts_desc" gorm:"comment:类别"`
	Series                        string    `json:"series" gorm:"comment:类别"`
	Pass_Date                     time.Time `json:"deferpass_datered_date" gorm:"comment:类别"`
	A3_Principal                  string    `json:"a3_principal" gorm:"comment:类别"`
	A3_Team_Top                   string    `json:"a3_team_top" gorm:"comment:类别"`
	A3_Team_Member                string    `json:"a3_team_member" gorm:"comment:类别"`
	A3_Issue_Desc                 string    `json:"a3_issue_desc" gorm:"comment:类别"`
	A3_Contingency                string    `json:"a3_contingency" gorm:"comment:类别"`
	A3_Cause                      string    `json:"a3_cause" gorm:"comment:类别"`
	A3_Affirm                     string    `json:"a3_affirm" gorm:"comment:类别"`
	A3_Correct                    string    `json:"a3_correct" gorm:"comment:类别"`
	A3_Verify                     string    `json:"a3_verify" gorm:"comment:类别"`
	A3_Result                     string    `json:"a3_result" gorm:"comment:类别"`
	A3_End                        string    `json:"a3_end" gorm:"comment:类别"`
	// 类别
}

type SetStatus struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

type SetPassDatte struct {
	ID             int64  `json:"id"`
	Operation_Type string `json:"operation_type" gorm:"comment:类别"` //1 返工 2 返修  3 配做 4 让步放行
}

func (Manage) TableName() string {
	return "ncr_product_manage"
}

type Complaint struct {
	global.GVA_MODEL
	Product_Name           string    `json:"product_name" gorm:"comment:类别"`
	Product_Sequence       string    `json:"product_sequence" gorm:"comment:类别"`
	Client_Name            string    `json:"client_name" gorm:"comment:类别"`
	Complaint_Name         string    `json:"complaint_name" gorm:"comment:类别"`
	Project_Name           string    `json:"project_name" gorm:"comment:类别"`
	Status                 string    `json:"Status" gorm:"comment:类别"`
	Checkout_Number        string    `json:"checkout_number" gorm:"comment:类别"`
	Interior_Feedback_Name string    `json:"interior_feedback_name" gorm:"comment:类别"`
	Interior_Feedback_Unit string    `json:"interior_feedback_unit" gorm:"comment:类别"`
	Issue_Desc             string    `json:"issue_desc" gorm:"comment:类别"`
	Issue_Level            string    `json:"issue_level" gorm:"comment:类别"`
	Short_Plan_Date        time.Time `json:"short_plan_date" gorm:"comment:类别"`
	Cause_Desc             string    `json:"cause_desc" gorm:"comment:类别"`
	Complaint_Order        string    `json:"complaint_order" gorm:"comment:类别"`
	Principal_Name         string    `json:"principal_name" gorm:"comment:类别"`
	Cost                   float64   `json:"Cost" gorm:"comment:类别"`
	Rectify_Date           time.Time `json:"rectify_date" gorm:"comment:类别"`
	Submit_Date            time.Time `json:"submit_date" gorm:"comment:类别"`
	Close_Date             time.Time `json:"close_date" gorm:"comment:类别"`
	Rectify                string    `json:"rectify " gorm:"comment:类别"`
}

func (Complaint) TableName() string {
	return "ncr_complaint"
}
