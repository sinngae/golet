package internal

import (
	"github.com/sinngae/golet/cmd/golet/internal/gin"
	"github.com/spf13/cobra"
)

var (
	cmdCli = &cobra.Command{
		Use:   "cli",
		Short: "start a http client",
		Run: func(cmd *cobra.Command, args []string) {
			gin.Startup(startParam)
		},
	}
)

func init() {
	cmdCli.PersistentFlags().StringVarP(&startParam, "start", "s", "port:8080", "Gin Server start @port")
	cmdRoot.AddCommand(cmdCli)
}
