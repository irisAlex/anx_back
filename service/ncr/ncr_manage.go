package ncr

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	"gorm.io/gorm"
)

type ManageService struct{}

func (apiService *ManageService) CreateManage(supplier ncr.Manage) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *ManageService) DeleteManage(supplier ncr.Manage) (err error) {
	var entity ncr.Manage
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

func (apiService *ManageService) GetManageInfoList(api ncr.Manage, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ncr.Manage{})
	var apiList []ncr.Manage

	if api.Department != "" {
		db = db.Where("department = ?", api.Department)
	}
	if api.Type_M != "" {
		db = db.Where("type_m = ?", api.Type_M)
	}

	if api.Category != "" {
		db = db.Where("category = ?", api.Category)
	}

	if api.Checkout_Name != "" {
		db = db.Where("checkout_name = ?", api.Checkout_Name)
	}
	if api.Checkout_Number != 0 {
		db = db.Where("checkout_number = ?", api.Checkout_Number)
	}

	if api.Project != "" {
		db = db.Where("project = ?", api.Project)
	}

	if api.Process_Mode != "" {
		db = db.Where("process_mode = ?", api.Process_Mode)
	}

	if !api.CreatedAt.IsZero() {
		db = db.Where(" created_at = ? ", api.CreatedAt.Format("2006-01-02"))
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

func (apiService *ManageService) GetManageById(id int) (api ncr.Manage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *ManageService) UpdateManage(api ncr.Manage) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}
