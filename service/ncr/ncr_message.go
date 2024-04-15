package ncr

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	"gorm.io/gorm"
)

type MessageService struct{}

func (apiService *MessageService) CreateMessage(supplier ncr.Message) (err error) {
	return global.GVA_DB.Create(&supplier).Error
}

func (apiService *MessageService) DeleteMessage(supplier ncr.Message) (err error) {
	var entity ncr.Message
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

func (apiService *MessageService) GetMessageInfoList(api ncr.Message, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&ncr.Message{})
	var apiList []ncr.Message

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

func (apiService *MessageService) GetMessageByname(name string) (list interface{}, err error) {

	fmt.Println(name)
	fmt.Println(3333)
	var apiList []ncr.Message
	err = global.GVA_DB.Model(&ncr.Message{}).Where("message_receive_name = ?", name).Find(&apiList).Error
	fmt.Println(4444)
	fmt.Println(apiList)
	return apiList, err
}

func (apiService *MessageService) UpdateMessage(api ncr.Message) (err error) {
	err = global.GVA_DB.Save(&api).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *MessageService) SetNcrStatus(ID uint, st string) (err error) {
	err = global.GVA_DB.Model(&ncr.Message{}).Where("id = ?", ID).Update("status", st).Error
	return err
}
