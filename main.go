package main

import (
	"demoproject/conf"
	_ "demoproject/docs"
	"demoproject/route"
	"demoproject/util"
)

// @title a demo
// @version 1.0
// @description A demo
// @termsOfService https://hello.com
// @license.name Apache 2.0
// @license.url https://hi.com
// @host localhost:8888
// @BasePath

func main() {
	conf.InitFlag()
	util.InitLog()

	util.Logzap.Info("Roooooooot...")
	//runtime.GOMAXPROCS()
	conf.Setup()
	go util.CronJob()       //定时任务
	go util.MqReceive(conf.MqCh,"withdraw")

	//gin.SetMode(gin.ReleaseMode)
	r := route.NewRoute()
	r.Run(":8888")

}
