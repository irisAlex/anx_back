package ncr

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	ncreq "github.com/flipped-aurora/gin-vue-admin/server/model/ncr/request"
	ncrep "github.com/flipped-aurora/gin-vue-admin/server/model/ncr/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type ComplaintApi struct{}

func (s *ComplaintApi) CreateComplaint(c *gin.Context) {
	var Complaint ncr.Complaint
	err := c.ShouldBindJSON(&Complaint)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println(Complaint)

	err = utils.Verify(Complaint, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ComplaintService.CreateComplaint(Complaint)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *ComplaintApi) GetComplaintList(c *gin.Context) {
	var pageInfo ncreq.SearchComplaintParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := ComplaintService.GetComplaintInfoList(pageInfo.Complaint, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (s *ComplaintApi) GetComplaintById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	api, err := ComplaintService.GetComplaintById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(ncrep.ComplaintResponse{Ncr: api}, "获取成功", c)
}

func (s *ComplaintApi) DeleteComplaint(c *gin.Context) {
	var api ncr.Complaint
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api.GVA_MODEL, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ComplaintService.DeleteComplaint(api)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *ComplaintApi) UpdateComplaint(c *gin.Context) {
	var api ncr.Complaint
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ComplaintService.UpdateComplaint(api)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (s *ComplaintApi) SetStatus(c *gin.Context) {
	var api ncr.SetStatus
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ComplaintService.SetNcrStatus(uint(api.ID), api.Status)
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}
