package control

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/config"
)

// CreateWorksheet 创建课表模板
func CreateWorksheet(f *excelize.File, weeks int) {
	// 创建一个工作表
	for i := 1; i <= weeks; i++ {
		f.NewSheet(fmt.Sprint("week", i))
		fmt.Println(i)
		// 设置单元格的值
		f.SetCellValue(fmt.Sprint("week", i), "A1", "节次/星期")
		for index := 1; index <= 13; index++ {
			f.SetCellValue(fmt.Sprint("week", i), fmt.Sprint("A", index+1), fmt.Sprint("第", index, "节"))
		}
		for index2, value := range config.WeekPosition2 {
			f.SetCellValue(fmt.Sprint("week", i), fmt.Sprint(value, "1"), config.Week[index2])
		}
	}

	// 根据指定路径保存文件
	if err := f.SaveAs(config.FileName); err != nil {
		fmt.Println(err)
	}
}
