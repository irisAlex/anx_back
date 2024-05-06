package excel

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	_ "image/jpeg"
	_ "image/png"

	"github.com/flipped-aurora/gin-vue-admin/server/model/ncr"
	"github.com/xuri/excelize/v2"
)

type Fruit struct {
	Id    uint
	Name  string
	Price float64
}

type Style struct {
	Border        []Border    `json:"border"`
	Fill          Fill        `json:"fill"`
	Font          *Font       `json:"font"`
	Alignment     *Alignment  `json:"alignment"`
	Protection    *Protection `json:"protection"`
	NumFmt        int         `json:"number_format"`
	DecimalPlaces int         `json:"decimal_places"`
	CustomNumFmt  *string     `json:"custom_number_format"`
	Lang          string      `json:"lang"`
	NegRed        bool        `json:"negred"`
}

// Border 边框
type Border struct {
	Type  string `json:"type"`
	Color string `json:"color"`
	Style int    `json:"style"`
}

// Fill 填充
type Fill struct {
	Type    string   `json:"type"`
	Pattern int      `json:"pattern"`
	Color   []string `json:"color"`
	Shading int      `json:"shading"`
}

// Font 字体
type Font struct {
	Bold      bool    `json:"bold"`      // 是否加粗
	Italic    bool    `json:"italic"`    // 是否倾斜
	Underline string  `json:"underline"` // single    double
	Family    string  `json:"family"`    // 字体样式
	Size      float64 `json:"size"`      // 字体大小
	Strike    bool    `json:"strike"`    // 删除线
	Color     string  `json:"color"`     // 字体颜色
}

// Protection 保护
type Protection struct {
	Hidden bool `json:"hidden"`
	Locked bool `json:"locked"`
}

// Alignment 对齐
type Alignment struct {
	Horizontal      string `json:"horizontal"`        // 水平对齐方式
	Indent          int    `json:"indent"`            // 缩进  只要设置了值，就变成了左对齐
	JustifyLastLine bool   `json:"justify_last_line"` // 两端分散对齐，只有在水平对齐选择 distributed 时起作用
	ReadingOrder    uint64 `json:"reading_order"`     // 文字方向 不知道值范围和具体的含义
	RelativeIndent  int    `json:"relative_indent"`   // 不知道具体的含义
	ShrinkToFit     bool   `json:"shrink_to_fit"`     // 缩小字体填充
	TextRotation    int    `json:"text_rotation"`     // 文本旋转
	Vertical        string `json:"vertical"`          // 垂直对齐
	WrapText        bool   `json:"wrap_text"`         // 自动换行
}

var Sheet = "Sheet1"

type issueDesc struct {
	Describe string `json:"describe" gorm:"comment:类别"`
	Issue    string `json:"issue" gorm:"comment:类别"`
	Isissue  bool   `json:"isIssue" gorm:"comment:类别"`
}

type commIssue struct {
	Name   string    `json:"name" gorm:"comment:类别"`
	Date   time.Time `json:"date" gorm:"comment:类别"`
	Scheme string    `json:"scheme" gorm:"comment:类别"`
	State  int       `json:"state" gorm:"comment:类别"`
}

type Issues struct {
	Name string `json:"name" gorm:"comment:类别"`
}

func CreateFile(m ncr.Manage) error {
	var f = excelize.NewFile()
	index, _ := f.NewSheet(Sheet)
	f.SetActiveSheet(index)

	x := 1
	y := 1
	black := "#000000"
	ControlsWriteExcel(f, "A2", "AD12", "Problem solving sheet 问题解决单A3报告     ( F-QA-0006-0038 )  ", SetStyle(f,
		40, "Frutiger LT 55 Roman", black, "left", "center", false))
	ControlsWriteExcel(f, "AE2", "AL12", "组长:"+m.A3_Team_Top, SetStyle(f,
		18, "宋体", black, "left", "center", false))
	ControlsWriteExcel(f, "AM2", "AU12", "组员:"+m.A3_Team_Member, SetStyle(f,
		18, "宋体", black, "left", "center", false))
	//问题1
	ControlsWriteExcel(f, "A13", "AU17", "1."+"Finding of facts / team-building 探究事实/团队合", SetStyle(f,
		22, "Frutiger LT 55 Roman", black, "left", "center", true))

	ControlsWriteExcel(f, "A18", "AU22", "Problem 问题描述:"+m.Defect_Problem, SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", false))

	ControlsWriteExcel(f, "V23", "AU64", m.Describe, SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", false))
	x = 23
	y = 28
	cellName := ""
	cellName2 := ""
	for i := 0; i < 7; i++ {
		switch i {
		case 0:
			cellName = "Location 车间"
			cellName2 = ""
		case 1:
			cellName = "Area/Work station工序/工位"
			cellName2 = m.Find_Process
		case 2:
			cellName = "Product/ Engineering part 产品部件号"
			cellName2 = m.Serialnumber
		case 3:
			cellName = "Defects per object 缺陷"
			cellName2 = m.Defect_Problem
		case 4:
			cellName = "discovered at被发现在"
			cellName2 = ""
		case 5:
			cellName = "discovered from被发现从"
			cellName2 = ""
		case 6:
			cellName = "recurring 重复problem  问题？"
			cellName2 = "是"
		}

		ControlsWriteExcel(f, "A"+intTostr(x), "J"+intTostr(y), cellName, SetStyle(f,
			20, "Frutiger LT 55 Roman", black, "left", "center", false))
		ControlsWriteExcel(f, "K"+intTostr(x), "U"+intTostr(y), cellName2, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		x += 6
		y += 6
	}

	//问题2
	ControlsWriteExcel(f, "A65", "AU70", "2. Problem description / symptoms问题描述", SetStyle(f,
		22, "Frutiger LT 55 Roman", black, "left", "center", true))
	ControlsWriteExcel(f, "A71", "P76", "描述", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", true))
	ControlsWriteExcel(f, "Q71", "AF76", "the problem is (NOK product / process) 这个问题是（不合格的产品/过程）", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", true))
	ControlsWriteExcel(f, "AG71", "AU76", `the problem is not (comparison with a OK product / process） 在其他产品/过程）就不是问题了`, SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))

	x = 77
	y = 82

	i2 := []issueDesc{}
	json.Unmarshal([]byte(m.A3_Issue_Desc), &i2)
	w5 := ""

	for key, val := range i2 {
		switch key {
		case 0:
			w5 = "What"
		case 1:
			w5 = "Where"
		case 2:
			w5 = "When"
		case 3:
			w5 = "Why"
		case 4:
			w5 = "Who"
		case 5:
			w5 = "How"
		case 6:
			w5 = "How Much/Many"
		}
		ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), w5, SetStyle(f,
			17, "Frutiger LT 55 Roman", black, "left", "center", false))
		ControlsWriteExcel(f, "E"+intTostr(x), "P"+intTostr(y), val.Describe, SetStyle(f,
			17, "Frutiger LT 55 Roman", black, "left", "center", false))
		ControlsWriteExcel(f, "Q"+intTostr(x), "AF"+intTostr(y), val.Issue, SetStyle(f,
			18, "宋体", black, "left", "center", false))
		ControlsWriteExcel(f, "AG"+intTostr(x), "AU"+intTostr(y), boolTostr(val.Isissue), SetStyle(f,
			18, "宋体", black, "center", "center", false))
		x += 6
		y += 6
	}
	//问题3
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "3. ERA", SetStyle(f,
		22, "Frutiger LT 55 Roman", black, "left", "center", true))
	x += 6
	y += 6
	ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), "Nr.", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "E"+intTostr(x), "U"+intTostr(y), "Emergency Response Action(s)应急响应措施（围堵）", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "V"+intTostr(x), "AJ"+intTostr(y), "Resp. 负责人", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "AH"+intTostr(x), "AP"+intTostr(y), "Date日期", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "AQ"+intTostr(x), "AU"+intTostr(y), "Status状态", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	x += 6
	y += 6

	i3 := []commIssue{}
	json.Unmarshal([]byte(m.A3_Contingency), &i3)
	for key, val := range i3 {
		ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), key+1, SetStyle(f,
			20, "Frutiger LT 55 Roman", black, "center", "center", false))
		ControlsWriteExcel(f, "E"+intTostr(x), "U"+intTostr(y), val.Name, SetStyle(f,
			18, "宋体", black, "left", "center", false))
		ControlsWriteExcel(f, "V"+intTostr(x), "AJ"+intTostr(y), val.Scheme, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		ControlsWriteExcel(f, "AH"+intTostr(x), "AP"+intTostr(y), val.Date.Format("2006-01-02"), SetStyle(f,
			18, "宋体", black, "center", "center", false))
		ControlsWriteExcelPicture(f, "AQ"+intTostr(x), "AU"+intTostr(y), val.State, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		x += 6
		y += 6
	}

	//问题4
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "Root cause analysis Ishikawa 原因分析鱼骨图", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", true))
	x += 24
	y += 24
	// 去掉 Base64 编码的头部
	imgBase64 := strings.Replace(m.A3_Img_Base64, "data:image/png;base64,", "", 1)

	// 解码 Base64 字符串
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		log.Fatal(err)
	}

	// // 将字节数组转为图片
	// img, _, err := image.Decode(strings.NewReader(string(imgBytes)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	InsertImg(f, "A"+intTostr(x), "AU"+intTostr(y), SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", false), imgBytes)
	x += 6
	y += 6
	//问题5
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "Root cause analysis (5 x why)根本原因确认(5个WHY)", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "left", "center", true))
	x += 6
	y += 6

	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "5 x why - apply the most probable root cause (s) from the Ishikawa 5个为什么-与鱼骨图合用确认问题根源", SetStyle(f,
		17, "Frutiger LT 55 Roman", black, "center", "center", true))
	x += 6
	y += 6
	var i4 []map[string]string
	json.Unmarshal([]byte(m.A3_Cause), &i4)
	issues := []Issues{}
	json.Unmarshal([]byte(m.A3_Issues), &issues)

	for key, val := range issues {
		var isList map[string]string
		ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), key+1, SetStyle(f,
			20, "Frutiger LT 55 Roman", black, "center", "center", true))
		ControlsWriteExcel(f, "E"+intTostr(x), "AU"+intTostr(y), val.Name, SetStyle(f,
			22, "宋体", black, "center", "center", false))
		x += 6
		y += 6
		if value, ok := getValue(i4, key); ok {
			isList = value
		}

		for i := 0; i < 5; i++ {
			ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), "why?", SetStyle(f,
				16, "Frutiger LT 55 Roman", black, "center", "center", false))
			if val, ok := isList[intTostr(i)]; ok {
				ControlsWriteExcel(f, "E"+intTostr(x), "AU"+intTostr(y), val, SetStyle(f,
					16, "宋体", black, "left", "center", false))
			} else {
				ControlsWriteExcel(f, "E"+intTostr(x), "AU"+intTostr(y), "无", SetStyle(f,
					16, "宋体", black, "left", "center", false))
			}
			x += 6
			y += 6
		}
	}
	//问题6
	x, y = batchWriteExectl(f, "Corrective action(s)纠正措施", "measures to remedy the identified cause (s)纠正措施过程", "chosen permanent corrective action(s)确定纠正措施:", m.A3_Correct, x, y)
	//问题7
	x, y = batchWriteExectl(f, "Effective措施验证     .", "Verification / Validation 检验/验证:", "(Dummy testing, trials, short-term capability indicators, audit, etc.假设检验，追踪，短期能力指数，审核等)", m.A3_Verify, x, y)
	//问题8
	x, y = batchWriteExectl(f, "Sustainability 有效地维持措施结果", "Sustainability of the successful solution (FMEA, control plan, instructions, training, etc.)", "保持改进结果 （FMEA，控制计划，zhi'dao书，培训，等）", m.A3_Result, x, y)
	//问题6
	x, y = batchWriteExectl(f, "Transfer传递", "Transfer of the solution to other products / processes (Lessons Learned) 把纠正措施应用到其它适用的产品/过程（经验分享）", "Proposal / Activity  标的/活动", m.A3_End, x, y)
	//foot
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "Com-pletion批准", SetStyle(f,
		22, "Frutiger LT 55 Roman", black, "left", "center", true))
	x += 6
	y += 6
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), "The Team-Leader presents the PSP to the Steering Committee.活动的组长提交给评审委员会     The confirmation of the Steering Committee relieved the Team-Leader.报告批准后，活动组长卸任", SetStyle(f,
		18, "宋体", black, "left", "center", false))
	x += 6
	y += 6
	ControlsWriteExcel(f, "A"+intTostr(x), "Y"+intTostr(y), "Team-Leader组长:", SetStyle(f,
		18, "宋体", black, "left", "center", false))
	ControlsWriteExcel(f, "Z"+intTostr(x), "AU"+intTostr(y), "	Completed on 批准日期:", SetStyle(f,
		18, "宋体", black, "left", "center", false))

	// 保存Excel
	if err := f.SaveAs("uploads/file/" + m.Serialnumber + "_A3.xlsx"); err != nil {
		return err
	}
	return nil
}

func intTostr(i int) string {
	return strconv.Itoa(i)
}

func ControlsWriteExcel(f *excelize.File, startCell, endCell string, content interface{}, styleID int) {
	f.MergeCell(Sheet, startCell, endCell)
	f.SetCellValue(Sheet, startCell, content)
	f.SetCellStyle(Sheet, startCell, endCell, styleID) //组长
}

func ControlsWriteExcelPicture(f *excelize.File, startCell, endCell string, state int, styleID int) {
	imgDir := "state4.png"
	f.MergeCell(Sheet, startCell, endCell)
	switch state {
	case 1:
		imgDir = "state4.png"
	case 2:
		imgDir = "state3.png"
	case 3:
		imgDir = "state2.png"
	case 4:
		imgDir = "state1.png"
	}
	f.AddPicture(Sheet, startCell, imgDir, &excelize.GraphicOptions{
		OffsetX:     60,
		OffsetY:     10,
		Positioning: "twoCell",
	})

	f.SetCellStyle(Sheet, startCell, endCell, styleID) //组长
}

func boolTostr(b bool) string {
	if b {
		return "有问题"
	}
	return "没有问题"
}

func getValue(slice []map[string]string, index int) (map[string]string, bool) {
	if index >= 0 && index < len(slice) {
		return slice[index], true
	}
	return nil, false
}

func batchWriteExectl(f *excelize.File, hTitle, sTitle, cTitle, fillContent string, x, y int) (int, int) {
	black := "#000000"
	ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), hTitle, SetStyle(f, 22, "Frutiger LT 55 Roman", black, "left", "center", true))
	x += 6
	y += 6
	if sTitle != "" {
		ControlsWriteExcel(f, "A"+intTostr(x), "AU"+intTostr(y), sTitle, SetStyle(f,
			22, "Frutiger LT 55 Roman", black, "center", "center", true))
	}
	x += 6
	y += 6
	ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), "Nr.", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "E"+intTostr(x), "U"+intTostr(y), cTitle, SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "V"+intTostr(x), "AJ"+intTostr(y), "Resp. 负责人", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "AH"+intTostr(x), "AP"+intTostr(y), "Date日期", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	ControlsWriteExcel(f, "AQ"+intTostr(x), "AU"+intTostr(y), "Status状态", SetStyle(f,
		20, "Frutiger LT 55 Roman", black, "center", "center", true))
	x += 6
	y += 6
	ii := []commIssue{}
	json.Unmarshal([]byte(fillContent), &ii)
	for key, val := range ii {
		ControlsWriteExcel(f, "A"+intTostr(x), "D"+intTostr(y), key+1, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		ControlsWriteExcel(f, "E"+intTostr(x), "U"+intTostr(y), val.Name, SetStyle(f,
			18, "宋体", black, "left", "center", false))
		ControlsWriteExcel(f, "V"+intTostr(x), "AJ"+intTostr(y), val.Scheme, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		ControlsWriteExcel(f, "AH"+intTostr(x), "AP"+intTostr(y), val.Date.Format("2006-01-02"), SetStyle(f,
			18, "宋体", black, "center", "center", false))
		ControlsWriteExcelPicture(f, "AQ"+intTostr(x), "AU"+intTostr(y), val.State, SetStyle(f,
			18, "宋体", black, "center", "center", false))
		x += 6
		y += 6
	}
	return x, y
}

func SetStyle(f *excelize.File, size int, font string, color string, horizontal, vertical string, isFill bool) int {
	var fill = excelize.Fill{}
	if isFill {
		fill = excelize.Fill{
			Type:    "pattern",
			Color:   []string{"D3D3D3"},
			Pattern: 1,
		}
	}
	var withoutFillStyle = &excelize.Style{
		Border: []excelize.Border{{Type: "left", Color: "#000000", Style: 1}, {Type: "top", Color: "#000000", Style: 1}, {Type: "right", Color: "#000000", Style: 1}, {Type: "bottom", Color: "#000000", Style: 1}},
		Font: &excelize.Font{
			Bold:   true,
			Size:   float64(size),
			Family: font,
			Color:  color,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      horizontal,
			Indent:          1,
			JustifyLastLine: true,
			ReadingOrder:    2,
			RelativeIndent:  1,
			ShrinkToFit:     true,
			TextRotation:    0,
			Vertical:        vertical,
			WrapText:        true,
		},
		Fill: fill,
	}

	withSty, _ := f.NewStyle(withoutFillStyle)
	return withSty
}

func InsertImg(f *excelize.File, startCell, endCell string, styleID int, imgBase4 []byte) {
	f.MergeCell(Sheet, startCell, endCell)
	err := f.AddPictureFromBytes("Sheet1", startCell, &excelize.Picture{
		Extension: ".png",
		File:      imgBase4,
	})
	if err != nil {
		fmt.Println(err)
	}
	f.SetCellStyle(Sheet, startCell, endCell, styleID) //组长
}
