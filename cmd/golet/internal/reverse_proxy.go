package internal

import (
	"github.com/sinngae/golet/cmd/internal/gin"
	"github.com/spf13/cobra"
)

var (
	reverseProxyGin = &cobra.Command{
		Use:   "gin",
		Short: "start a Reverse Proxy Server",
		Run: func(cmd *cobra.Command, args []string) {
			gin.Startup(startParam)
		},
	}
)

func init() {
	reverseProxyGin.PersistentFlags().StringVarP(&startParam, "start", "s", "port=:8080", "Gin Server start @port")
	cmdRoot.AddCommand(reverseProxyGin)
}
