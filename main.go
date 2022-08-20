package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-CourseSchedule/api"
	"go-CourseSchedule/config"
)

func main() {
	// gin
	r := gin.Default()
	//获取当前课程
	r.GET("/getTodayCurriculum", api.GetTodayCurriculum)
	r.GET("/getNextClass", api.GetNextClass)
	err := r.Run(config.Port)
	if err != nil {
		fmt.Println(err)
	}
}
