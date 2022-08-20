package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-CourseSchedule/config"
	"go-CourseSchedule/control"
	"net/http"
)

// GetTodayCurriculum 获取当天课程
func GetTodayCurriculum(context *gin.Context) {
	xlsx := control.OpenXlsx()
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
}
