package internal

import (
	"fmt"
	"log"

	"github.com/sinngae/gland/cmd/gland/internal/json2xls"

	"github.com/spf13/cobra"
)

var (
	cmdDepart = &cobra.Command{
		Use:   "json2xml",
		Short: "trans json str to xml file",
		Run: func(cmd *cobra.Command, args []string) {
			err := json2xls.Json2Xls("", "")
			if err != nil {
				log.Fatalf("err=%s", err)
			}
			fmt.Printf("%s Done\n", cmd.Use)
		},
	}

	shareCount uint32
)

func init() {
	cmdDepart.PersistentFlags().Uint32VarP(&shareCount, "share", "s", 256, "table shares quantity to depart")
	cmdRoot.AddCommand(cmdDepart)
}
