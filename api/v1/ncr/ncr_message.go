package ncr

import (
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

type MessageApi struct{}

func (s *MessageApi) SendMessage(c *gin.Context) {
	var Message ncr.Message
	err := c.ShouldBindJSON(&Message)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(Message, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = MessageService.CreateMessage(Message)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *MessageApi) GetMessageList(c *gin.Context) {
	var pageInfo ncreq.SearchMessageParams
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
	list, total, err := MessageService.GetMessageInfoList(pageInfo.Message, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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

func (s *MessageApi) GetMessageByName(c *gin.Context) {
	var idInfo request.GetByName
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
	api, total, err := MessageService.GetMessageByname(idInfo.Name)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     api,
		Total:    total,
		Page:     10,
		PageSize: 100,
	}, "获取成功", c)
}

func (s *MessageApi) DeleteMessage(c *gin.Context) {
	var api ncr.Message
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
	err = MessageService.DeleteMessage(api)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *MessageApi) GetMessageMessageByNcrID(c *gin.Context) {
	var api request.GetByNcrID
	err := c.ShouldBindJSON(&api)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(api, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	message, err := MessageService.GetMessageByNcrID(api.NcrID, api.ReportName)
	if err != nil {
		global.GVA_LOG.Error("当前用户没有操作权限!", zap.Error(err))
		response.FailWithMessage("当前用户没有操作权限", c)
		return
	}

	response.OkWithDetailed(ncrep.MessageResponse{Ncr: message}, "获取成功", c)
}

func (s *MessageApi) UpdateMessage(c *gin.Context) {
	var api ncr.Message
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
	err = MessageService.UpdateMessage(api)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (s *MessageApi) SetStatus(c *gin.Context) {
	var api request.GetById
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
	err = MessageService.SetMessageState(uint(api.ID))
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}
