package cmd

import (
	"dbland/config"
	"dbland/repository"
	"dbland/router"
	"log/slog"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	port = 0

	startCmd = &cobra.Command{
		Use:   "start",
		Short: "Start the server",
		Long:  "Start the running web server gracefully.",
		Run: func(cmd *cobra.Command, args []string) {
			start()
		},
	}
)

func start() {
	slog.Info("Initialize config")
	config.InitConfig()

	slog.Info("Initialize router")
	r := router.InitRouter()

	slog.Info("Initialize db")
	repository.Initialize()

	slog.Info("Start server", "port", port)
	var err error
	go func() {
		err = r.Run(":" + strconv.Itoa(port))
		if err != nil {
			slog.Error(err.Error())
			return
		} else {
			slog.Info("Service started successfully!")
		}
	}()
	select {}
}

func init() {
	RootCmd.AddCommand(startCmd)
}
