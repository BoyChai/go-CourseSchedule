# go-CourseSchedule

## 概述

北理课表规范化生成并提供api接口

## 工具

**xlsxtool：**用于生成每周的课表

**go-CourseSchedule：**用于提供api接口

## 配置文件

**Port：**api程序监听的接口

**FileName：**xlsx课表文件名称(格式需要和本项目中的格式一样，否则程序会出问题)

**Sheet：**xlsx工作表名称

**WeekPosition1：**xlsx课表的星期坐标字母

**WeekPosition2：**xlsx课表生成的星期坐标字母

**Week：**生成的星期标识名称，可以自行更改

**DataPosition：**xlsx课表的坐标

**FirstWeek：**课程的第一个周的时间(星期一)

**ClassHour：**每一节课的时间

## 使用

**xlsxtool**用于生成每周的课表格式分为两种，一种是只会生成课程名称和位置，另外一种是给api程序提供的。使用之前需要确包配置文件正确。

**xlsxtool**直接运行会生成一种已课程名称和位置的类型，执行`./xlsxtool -ty gin`会生成给api程序查看的格式。

当**xlsxtool**生成的格式是提供给api程序的，就可以开始运行api程序(`go-CourseSchedule`)了

## 接口

/getNextClass 			Get请求，返回当前课程或者下一节课的课程信息

/getTodayCurriculum 	Get请求，返回今天的全部课程