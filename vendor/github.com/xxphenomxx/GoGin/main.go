package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/xxphenomxx/GoGin/pkg/setting"
	"github.com/xxphenomxx/GoGin/routers"
	"github.com/xxphenomxx/GoGin/models"
	"github.com/xxphenomxx/GoGin/pkg/logging"
	"github.com/xxphenomxx/GoGin/pkg/gredis"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()
}

// @title Gogin API
// @version 1.0
// @description Golang Gin API Boilerplate
// @termsOfService https://github.com/xxphenomxx/gogin
// @license.name MIT
// @license.url https://github.com/xxphenomxx/gogin/blob/master/LICENSE
func main() {
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)

	server.ListenAndServe()

	// If you want Graceful Restart, you need a Unix system and download github.com/fvbock/endless
	//endless.DefaultReadTimeOut = readTimeout
	//endless.DefaultWriteTimeOut = writeTimeout
	//endless.DefaultMaxHeaderBytes = maxHeaderBytes
	//server := endless.NewServer(endPoint, routersInit)
	//server.BeforeBegin = func(add string) {
	//	log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	log.Printf("Server err: %v", err)
	//}
}