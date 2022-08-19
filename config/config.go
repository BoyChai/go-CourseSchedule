package config

type CourseInformation struct {
	// 课程id
	ID int64 `json:"id"`
	// 课程名称
	Name string `json:"name"`
	// 第几周
	Week []int `json:"week"`
	// 星期
	DayWeek int `json:"dayWeek"`
	// 节次
	Class []int `json:"class"`
	// 位置
	Position string `json:"position"`
}

type CurriculumInfo struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
}

var FileName = "timeTableForStu12.xlsx"
var WeekPosition1 = []string{"C", "D", "E", "F", "G", "H", "I"}
var WeekPosition2 = []string{"B", "C", "D", "E", "F", "G", "H"}
var Week = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期七"}
var DataPosition = []string{"14", "20", "26", "32", "38"}
var Sheet = "sheet1"
