package ncr

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	"gorm.io/gorm"
)

type ProjectService struct{}

func (apiService *ProjectService) CreateProject(supplier ncr.Project) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *ProjectService) DeleteProject(supplier ncr.Project) (err error) {
	var entity ncr.Project
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

func (apiService *ProjectService) GetProjectInfoList(api ncr.Project, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ncr.Project{})
	var apiList []ncr.Project

	if api.Name != "" {
		db = db.Where("name = ?", api.Name)
	}
	if api.Describe != "" {
		db = db.Where("describe = ?", api.Describe)
	}

	if api.Period > 0 {
		db = db.Where("period = ?", api.Period)
	}

	if api.Principal != "" {
		db = db.Where("principal = ?", api.Principal)
	}

	if api.Priority != "" {
		db = db.Where("priority = ?", api.Priority)
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
			orderMap["period"] = true
			orderMap["principal"] = true
			orderMap["describe"] = true
			orderMap["priority"] = true
			orderMap["name"] = true
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

func (apiService *ProjectService) GetProjectById(id int) (api ncr.Project, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&api).Error
	return
}

func (apiService *ProjectService) UpdateProject(api ncr.Project) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}
