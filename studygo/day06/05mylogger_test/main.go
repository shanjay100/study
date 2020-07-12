package main

import (

	"../mylogger"
	"time"
)

// 测试我们自己写的日志库
func main()  {
	//log := mylogger.NewLog("error")
	log := mylogger.NewFileLogger("Info", "./", "0712.log", 10*1024*1024)
	for {
		log.Debug("这是一条Debug日志")
		log.Info("这是一条Info日志")
		log.Warning("这是一条Warning日志")
		id := 10010
		name := "理想"
		log.Error("这是一条Error日志, id:%d, name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		time.Sleep(time.Second * 2)
	}
}
