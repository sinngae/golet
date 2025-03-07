package commands

import (
	"fmt"
	"log"

<<<<<<<< HEAD:cmd/servlet/commands/json2xls.go
	"github.com/sinngae/golet/cmd/servlet/src/handler/json2xls"
========
	"github.com/sinngae/golet/sample/cmdcli/golet/internal/json2xls"

>>>>>>>> 4aad3e1b64427d5ebafb07f037b140f7eb3a6511:sample/cmdcli/golet/internal/json2xls.go
	"github.com/spf13/cobra"
)

var (
	cmdJson2Xls = &cobra.Command{
		Use:   "json2xml",
		Short: "trans json str to xml file",
		Run: func(cmd *cobra.Command, args []string) {
			err := json2xls.Json2Xls(input, output)
			if err != nil {
				log.Fatalf("err=%s", err)
			}
			fmt.Printf("%s Done\n", cmd.Use)
		},
	}

	shareCount uint32
)

func init() {
	cmdJson2Xls.PersistentFlags().Uint32VarP(&shareCount, "share", "s", 256, "table shares quantity to depart")
	cmdRoot.AddCommand(cmdJson2Xls)
}
