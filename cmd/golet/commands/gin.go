package commands

import (
	"fmt"
	"log"

	"github.com/sinngae/golet/cmd/servlet/src/handler/ginlet"

	"github.com/spf13/cobra"
)

var (
	cmdGin = &cobra.Command{
		Use:   "gin",
		Short: "start gin server",
		Run: func(cmd *cobra.Command, args []string) {
			err := ginlet.Handle()
			if err != nil {
				log.Fatalf("err=%s", err)
			}
			fmt.Printf("%s Done\n", cmd.Use)
		},
	}

	//shareCount uint32
)

func init() {
	cmdGin.PersistentFlags().Uint32VarP(&shareCount, "share", "s", 256, "table shares quantity to depart")
	cmdRoot.AddCommand(cmdGin)
}
