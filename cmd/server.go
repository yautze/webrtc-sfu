package cmd

import (
	"webrtc-sfu/infra"
	"webrtc-sfu/infra/config"
	"webrtc-sfu/router"

	"github.com/kataras/iris/v12"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&config.C.Log.Level, "log level", "l", "debug", "server log level")
}

func start() {
	app := iris.Default()

	infra.Init()

	// router
	router.Set(app)

	app.Run(iris.Addr(":3000"), iris.WithoutInterruptHandler)
}
