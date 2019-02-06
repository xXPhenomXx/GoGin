package main

import (
	"GoGin/cmd"
)


// @title Gogin API
// @version 1.0
// @description Golang Gin API Boilerplate
// @termsOfService https://github.com/xxphenomxx/gogin
// @license.name MIT
// @license.url https://github.com/xxphenomxx/gogin/blob/master/LICENSE
func main() {

	cmd.Execute()

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
