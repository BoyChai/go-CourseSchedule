package control

import (
	"errors"
	"fmt"
	"go-CourseSchedule/config"
	"math"
	"strconv"
	"strings"
	"time"
)

// TimeSign 计算时间表示
// 代码写的很乱 求大佬优化
func TimeSign(ty string) (int, error) {
	//当前时间
	t1 := time.Now()
	//时间格式
	layout := []string{
		"2006-01-02",
		"15:04",
	} //周开始时间
	t2, err := time.Parse(layout[0], config.FirstWeek)
	if err != nil {
		fmt.Println(err)
	}
	switch ty {
	case "w":
		//时间标记(第几周)
		sign := int(math.Ceil(t1.Sub(t2).Hours() / 24 / 7))
		return sign, nil
	case "d":
		//时间标记(第几天)
		week := int(math.Floor(t1.Sub(t2).Hours()/24/7)) * 7
		sign := int(math.Ceil(t1.Sub(t2).Hours()/24)) - week
		return sign, nil
	case "h":
		//时间标记(当前课时或者下一节课)
		h1, err := strconv.Atoi(strings.Split(fmt.Sprint(t1.Format(layout[1])), ":")[0])
		if err != nil {
			fmt.Println(err)
		}
		for i, _ := range config.ClassHour {
			h2, err := strconv.Atoi(strings.Split(strings.Split(config.ClassHour[i], "-")[1], ":")[0])
			if err != nil {
				fmt.Println(err)
			}
			if h1 < h2 {
				return 1, nil
			}
			// 判断当前时是否和下课时间的时是否相等
			if h1 == h2 {
				// 判断当前分是否小于或等于下课时间的分
				if transformation(strings.Split(fmt.Sprint(t1.Format(layout[1])), ":")[1]) <= transformation(strings.Split(strings.Split(config.ClassHour[i], "-")[1], ":")[1]) {
					return i + 2, nil
				}
				// 判断当前时是否大于等于上课时间时
				if h1 >= transformation(strings.Split(strings.Split(config.ClassHour[i], "-")[0], ":")[0]) {
					return i + 1, nil
				}
				//判断当前分是否大于上课时间分
				if transformation(strings.Split(fmt.Sprint(t1.Format(layout[1])), ":")[1]) > transformation(strings.Split(strings.Split(config.ClassHour[i], "-")[0], ":")[1]) {
					return i + 1, nil
				}
			}
			return i + 1, nil
		}
	default:
		return 0, errors.New("时间类型错误")
	}
	return 0, nil
}

func transformation(s string) int {
	value, err := strconv.Atoi(s)
	fmt.Println(err)
	return value
}
