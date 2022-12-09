package internal

import (
	"github.com/sinngae/golet/cmd/golet/internal/gin"
	"github.com/spf13/cobra"
)

var (
	cmdGin = &cobra.Command{
		Use:   "gin",
		Short: "start a Gin server",
		Run: func(cmd *cobra.Command, args []string) {
			gin.Startup(startParam)
		},
	}

	startParam string
)

func init() {
	cmdGin.PersistentFlags().StringVarP(&startParam, "start", "s", "port=:8080", "Gin Server start @port")
	cmdRoot.AddCommand(cmdGin)
}
