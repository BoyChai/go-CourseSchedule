package main

import (
	"flag"
	"fmt"
	"go-CourseSchedule/config"
	"go-CourseSchedule/control"
)

func main() {
	config.ReadConfig()
	ty := flag.String("ty", "gin", "生成模板类型")
	flag.Parse()
	xlsx := control.OpenXlsx()
	// 格式化课表
	control.CreateWorksheet(xlsx, 19)
	// 获取课表信息
	allCell := control.ReadTimetable(xlsx)
	// 生成课表信息
	control.CourseClassification(xlsx, allCell, fmt.Sprint(ty))

}
