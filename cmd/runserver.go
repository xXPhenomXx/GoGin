package cmd

import (
	"fmt"
	"log"
	"net/http"
	"github.com/spf13/cobra"
	"GoGin/pkg/setting"
	"GoGin/routers"
	"GoGin/models"
	"GoGin/pkg/logging"
	"GoGin/pkg/gredis"
)

var runServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Starts GoGin server",
	Long:  `This subcommand says hello`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("GoGin server started")
	},
}

func init() {
	RootCmd.AddCommand(runServerCmd)
	setting.Setup()
	models.Setup()
	logging.Setup()
	gredis.Setup()

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
}
