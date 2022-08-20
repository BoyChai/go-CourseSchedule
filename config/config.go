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
	Class    int    `json:"class"`
	Position string `json:"position"`
}

type Request struct {
	Week           int              `json:"week"`
	DayWeek        int              `json:"dayWeek"`
	CurriculumInfo []CurriculumInfo `json:"curriculumInfo"`
}

//
//var FileName string
//var WeekPosition1 = []string{"C", "D", "E", "F", "G", "H", "I"}
//var WeekPosition2 = []string{"B", "C", "D", "E", "F", "G", "H"}
//var Week = []string{"星期一", "星期二", "星期三", "星期四", "星期五", "星期六", "星期日"}
//var DataPosition = []string{"14", "20", "26", "32", "38"}
//var Sheet = "sheet1"
//var FirstWeek = "2022-08-15"
//var ClassHour = []string{
//	"8:00-8:45",
//	"8:50-9:35",
//	"9:55-10:40",
//	"10:45-11:30",
//	"11:35-12:20",
//	"13:20-14:05",
//	"14:10-14:55",
//	"15:15-16:00",
//	"16:05-16:50",
//	"16:55-17:40",
//	"18:30-19:15",
//	"19:20-20:05",
//	"20:10-20:55",
//}
var Port string
var FileName string
var WeekPosition1 []string
var WeekPosition2 []string
var Week []string
var DataPosition []string
var Sheet string
var FirstWeek string
var ClassHour []string
