package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func ReadConfig() {
	v := viper.New()
	v.AddConfigPath("./config")
	v.SetConfigName("config")
	v.SetConfigType("ini")
	err := v.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	FileName = v.GetString("default.FileName")
	Sheet = v.GetString("default.Sheet")
	WeekPosition1 = strings.Split(v.GetString("default.WeekPosition1"), ",")
	WeekPosition2 = strings.Split(v.GetString("default.WeekPosition2"), ",")
	Week = strings.Split(v.GetString("default.Week"), ",")
	DataPosition = strings.Split(v.GetString("default.DataPosition"), ",")
	FirstWeek = v.GetString("default.FirstWeek")
	ClassHour = strings.Split(v.GetString("default.ClassHour"), ",")
}
