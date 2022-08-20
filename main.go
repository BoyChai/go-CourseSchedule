package main

import (
	"github.com/gin-gonic/gin"
	"go-CourseSchedule/api"
)

func main() {
	// gin
	r := gin.Default()
	//获取当前课程
	r.GET("/getTodayCurriculum", api.GetTodayCurriculum)
	r.GET("/getNextClass", api.GetNextClass)
	r.Run()
}
