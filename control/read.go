package control

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/config"
	"strings"
)

// ReadTimetable 读取全部的课程信息
func ReadTimetable(f *excelize.File) []string {
	var allCell []string
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 读取每星期的课
	for _, wpValue := range config.WeekPosition1 {
		//拿到坐标进行读出
		for _, dpValue := range config.DataPosition {
			cell, err := f.GetCellValue(config.Sheet, fmt.Sprintf(wpValue+dpValue))
			if err != nil {
				fmt.Println(err)
				return nil
			}
			// 读出之后切割，切割之后写入allCell这个切片
			cellValue := strings.Split(cell, "\n")
			for _, cv := range cellValue {
				allCell = append(allCell, cv)
			}
		}
	}
	return allCell
}
