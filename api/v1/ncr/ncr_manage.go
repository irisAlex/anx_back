package ncr

import (
	"fmt"
	"os"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	ncreq "github.com/flipped-aurora/gin-vue-admin/server/model/ncr/request"
	ncrep "github.com/flipped-aurora/gin-vue-admin/server/model/ncr/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/excel"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type ManageApi struct{}

func (s *ManageApi) GetFile(c *gin.Context) {
	f, _ := os.Open("A3.xlsx")
	defer f.Close()

	p := make([]byte, 1024)

	w := c.Writer
	w.Header().Set("Content-Type", "application/vnd.ms-excel")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", "A3.xlsx"))

	var readErr error
	var readCount int

	for {
		readCount, readErr = f.Read(p)
		if readErr != nil {
			break
		}
		if readCount > 0 {
			if _, err := w.Write(p[:readCount]); err != nil {
				break
			}
		}
	}
	w.Flush()
	//	c.File("A3.xlsx")
	response.Ok(c)
}

func (s *ManageApi) CreateManage(c *gin.Context) {
	var Manage ncr.Manage
	err := c.ShouldBindJSON(&Manage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println(Manage)

	err = utils.Verify(Manage, utils.ApiVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = ManageService.CreateManage(Manage)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

func (s *ManageApi) GetManageList(c *gin.Context) {
	var pageInfo ncreq.SearchManageParams
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
	list, total, err := ManageService.GetManageInfoList(pageInfo.Manage, pageInfo.PageInfo, pageInfo.OrderKey, pageInfo.Desc)
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

func (s *ManageApi) GetManageById(c *gin.Context) {
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
	api, err := ManageService.GetManageById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(ncrep.ManageResponse{Ncr: api}, "获取成功", c)
}

func (s *ManageApi) CloseAllById(c *gin.Context) {
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
	err = ManageService.CloseAllById(uint(idInfo.ID))
	if err != nil {
		global.GVA_LOG.Error("关闭失败!", zap.Error(err))
		response.FailWithMessage("关闭失败", c)
		return
	}
	response.Ok(c)
	//response("获取成功", c)
	//response.OkWithDetailed(ncrep.ManageResponse{Ncr: api}, "获取成功", c)
}

func (s *ManageApi) DeleteManage(c *gin.Context) {
	var api ncr.Manage
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
	err = ManageService.DeleteManage(api)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (s *ManageApi) UpdateManage(c *gin.Context) {
	var api ncr.Manage
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
	err = ManageService.UpdateManage(api)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

func (s *ManageApi) SetStatus(c *gin.Context) {
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
	err = ManageService.SetNcrStatus(uint(api.ID), api.Status)
	if err != nil {
		global.GVA_LOG.Error("更新状态失败!", zap.Error(err))
		response.FailWithMessage("更新状态失败", c)
		return
	}
	response.OkWithMessage("更新状态成功", c)
}

func (s *ManageApi) SetPassDate(c *gin.Context) {
	var api ncr.SetPassDatte
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
	err = ManageService.SetNcrPassDate(uint(api.ID), api.Operation_Type)
	if err != nil {
		global.GVA_LOG.Error("更新放行时间失败!", zap.Error(err))
		response.FailWithMessage("更新放行时间失败", c)
		return
	}
	response.OkWithMessage("更新放行时间成功", c)
}

func (s *ManageApi) SetNcr(c *gin.Context) {
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
	err = ManageService.SetNcr(uint(api.ID))
	if err != nil {
		global.GVA_LOG.Error("创建NCR失败!", zap.Error(err))
		response.FailWithMessage("创建NCR失败", c)
		return
	}
	response.OkWithMessage("创建NCR成功", c)
}

func (s *ManageApi) UpdateParts(c *gin.Context) {
	var api ncr.Manage
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
	err = ManageService.UpdateParts(api)
	if err != nil {
		global.GVA_LOG.Error("更新配做失败!", zap.Error(err))
		response.FailWithMessage("更新配做失败", c)
		return
	}
	response.OkWithMessage("更新配做成功", c)
}

func (s *ManageApi) Down(c *gin.Context) {
	var api ncr.Manage
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
	err = ManageService.UpdateParts(api)
	if err != nil {
		global.GVA_LOG.Error("更新配做失败!", zap.Error(err))
		response.FailWithMessage("更新配做失败", c)
		return
	}
	response.OkWithMessage("更新配做成功", c)
}

func (s *ManageApi) DownExcelFile(c *gin.Context) {
	var idInfo request.GetById
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id 为空", c)
		return
	}
	idInfo.ID = strToint(id)
	api, err := ManageService.GetManageById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	err = excel.CreateFile(api)
	if err != nil {
		global.GVA_LOG.Error("Excel 文件生成失败!", zap.Error(err))
		response.FailWithMessage("Excel 文件生成失败!", c)
		return
	}
	response.OkWithDetailed(ncrep.ManageResponse{Ncr: api}, "下载成功", c)
}

func strToint(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
