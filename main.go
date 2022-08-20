package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/config"
	"go-CourseSchedule/control"
	"net/http"
)

func main() {
	config.ReadConfig()
	xlsx, err := excelize.OpenFile(config.FileName)
	if err != nil {
		fmt.Println(err)
	}
	// 格式化课表
	control.CreateWorksheet(xlsx, 19)
	// 获取课表信息
	allCell := control.ReadTimetable(xlsx)
	// 生成课表信息
	control.CourseClassification(xlsx, allCell, "gin")

	// gin
	r := gin.Default()

	r.GET("/getTodayCurriculum", func(context *gin.Context) {
		var Request config.Request
		//获取周
		w, err := control.TimeSign("w")
		if err != nil {
			fmt.Println(err)
		}
		//获取天
		d, err := control.TimeSign("d")
		if err != nil {
			fmt.Println(err)
		}
		Request.Week = w
		Request.DayWeek = d
		for i := 0; i < len(config.ClassHour); i++ {
			value, err := xlsx.GetCellValue(fmt.Sprint("week", w), fmt.Sprint(config.WeekPosition2[d-1], i+2))
			if err != nil {
				return
			}
			var CurriculumInfo config.CurriculumInfo
			value = fmt.Sprintf(value)
			json.Unmarshal([]byte(value), &CurriculumInfo)
			Request.CurriculumInfo = append(Request.CurriculumInfo, CurriculumInfo)

		}
		//返回数据
		context.JSON(http.StatusOK, Request)
	})
	r.Run()
}
