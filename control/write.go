package control

import (
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/config"
	"regexp"
	"strconv"
	"strings"
)

func CourseClassification(f *excelize.File, allCell []string, ty string) {
	for index, value := range allCell {
		var CourseInfo config.CourseInformation
		// 判断课程规范语句
		match, err := regexp.MatchString("[0-9-]+周,星期[0-9],第[0-9]*节-第[0-9]*节", value)
		if err != nil {
			fmt.Println(err)
		}
		// 判断成功
		if match {
			cell := strings.Split(allCell[index-1], "-")
			//拿到id
			rs := []rune(cell[0])
			CourseInfo.ID, err = strconv.ParseInt(string(rs[3:]), 10, 64)
			//拿到课程名称
			CourseInfo.Name = cell[1]
			split := strings.Split(value, ",")
			// 拿到周的值
			weekSplit := strings.Split(split[0], "周")
			if len([]rune(weekSplit[0])) <= 2 {
				week, err := strconv.Atoi(weekSplit[0])
				if err != nil {
					fmt.Println(err)
				}
				CourseInfo.Week = append(CourseInfo.Week, week)
			} else {
				weekSplitSp := strings.Split(weekSplit[0], "-")
				wp1, err := strconv.Atoi(weekSplitSp[0])
				if err != nil {
					fmt.Println(err)
				}
				wp2, err := strconv.Atoi(weekSplitSp[1])
				if err != nil {
					fmt.Println(err)
				}
				for i := wp1; i <= wp2; i++ {
					CourseInfo.Week = append(CourseInfo.Week, i)
				}
			}
			// 拿到星期的值
			dayWeekSplit := strings.Split(split[1], "星期")
			dayWeek, err := strconv.Atoi(dayWeekSplit[1])
			if err != nil {
				fmt.Println(err)
			}
			CourseInfo.DayWeek = dayWeek
			// 拿到节的值
			classSplit := strings.Split(split[2], "-")
			rs1 := []rune(classSplit[0])
			classSplit[0] = strings.Split(string(rs1[1:]), "节")[0]
			rs1 = []rune(classSplit[1])
			classSplit[1] = strings.Split(string(rs1[1:]), "节")[0]
			classSplit1, err := strconv.Atoi(classSplit[0])
			if err != nil {
				fmt.Println(err)
			}
			classSplit2, err := strconv.Atoi(classSplit[1])
			if err != nil {
				fmt.Println(err)
			}
			for i := classSplit1; i <= classSplit2; i++ {
				CourseInfo.Class = append(CourseInfo.Class, i)
			}
			//拿到位置
			CourseInfo.Position = split[3]
			//info := fmt.Sprint("课程ID:", CourseInfo.ID, " 课程名称:", CourseInfo.Name, "周次:", CourseInfo.Week, "星期:", CourseInfo.DayWeek, "节次:", CourseInfo.Class, "位置:", CourseInfo.Position)
			//fmt.Println(CourseInfo)
			// json输出
			//jsonStr, err := json.Marshal(CourseInfo)
			//if err != nil {
			//	fmt.Println(err)
			//}
			//fmt.Println(string(jsonStr))
			write(f, CourseInfo, ty)
		}
	}

}

func write(f *excelize.File, CourseInfo config.CourseInformation, ty string) {
	for _, sheet := range CourseInfo.Week {
		for index, class := range CourseInfo.Class {
			var c = config.CurriculumInfo{
				ID:       CourseInfo.ID,
				Name:     CourseInfo.Name,
				Position: CourseInfo.Position,
				Class:    CourseInfo.Class[index],
			}
			jsonStr, err := json.Marshal(c)
			if err != nil {
				fmt.Println(err)
			}
			axis := fmt.Sprint(config.WeekPosition2[CourseInfo.DayWeek-1], class+1)
			//value := fmt.Sprint()
			if ty != "gin" {
				f.SetCellValue(fmt.Sprint("week", sheet), axis, fmt.Sprint(CourseInfo.Name, "\n", CourseInfo.Position))
			} else {
				f.SetCellValue(fmt.Sprint("week", sheet), axis, jsonStr)
			}
			f.Save()
		}
	}
}
