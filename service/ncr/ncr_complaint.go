package ncr

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	"gorm.io/gorm"
)

type ComplaintService struct{}

func (apiService *ComplaintService) CreateComplaint(supplier ncr.Complaint) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *ComplaintService) DeleteComplaint(supplier ncr.Complaint) (err error) {
	var entity ncr.Complaint
	err = global.GVA_DB.Where("id = ?", supplier.ID).First(&entity).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {                           // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	//CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ComplaintService) GetComplaintInfoList(api ncr.Complaint, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ncr.Complaint{})
	var apiList []ncr.Complaint

	if api.Interior_Feedback_Name != "" {
		db = db.Where("department = ?", api.Interior_Feedback_Name)
	}
	if api.Interior_Feedback_Unit != "" {
		db = db.Where("mold = ?", api.Interior_Feedback_Unit)
	}

	if api.Product_Sequence != "" {
		db = db.Where("category = ?", api.Product_Sequence)
	}

	if api.Product_Name != "" {
		db = db.Where("checkout_name = ?", api.Product_Name)
	}

	err = db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			// 设置有效排序key 防止sql注入
			// 感谢 Tom4t0 提交漏洞信息
			orderMap := make(map[string]bool, 5)
			orderMap["id"] = true
			if orderMap[order] {
				if desc {
					OrderStr = order + " desc"
				} else {
					OrderStr = order
				}
			} else { // didn't match any order key in `orderMap`
				err = fmt.Errorf("非法的排序字段: %v", order)
				return apiList, total, err
			}

			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Debug().Order("id").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

func (apiService *ComplaintService) GetComplaintById(id int) (api ncr.Complaint, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *ComplaintService) UpdateComplaint(api ncr.Complaint) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *ComplaintService) SetNcrStatus(ID uint, st string) (err error) {
	err = global.GVA_DB.Model(&ncr.Complaint{}).Where("id = ?", ID).Update("status", st).Error
	return err
}
