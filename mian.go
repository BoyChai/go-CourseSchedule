package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/control"
)

func main() {
	xlsx, err := excelize.OpenFile("timeTableForStu12.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	// 格式化课表
	control.CreateWorksheet(xlsx, 19)
	// 获取课表信息
	allCell := control.ReadTimetable(xlsx)
	// 生成课表信息
	control.CourseClassification(xlsx, allCell)

}
