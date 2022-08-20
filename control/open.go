package control

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"go-CourseSchedule/config"
)

func OpenXlsx() *excelize.File {
	config.ReadConfig()
	xlsx, err := excelize.OpenFile(config.FileName)
	if err != nil {
		fmt.Println(err)
	}
	return xlsx
}
