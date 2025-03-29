package commands

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	// flag value
	ctxCom context.Context
	input  string
	output string

	env string

	// cmdRoot
	cmdRoot = &cobra.Command{
		Use:   "golet",
		Short: "golet is a tool for Dev daily work",
		Long:  `golet is a tool for developer daily work.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctxCom = context.WithValue(context.Background(), "host", "zqren-test")
			ctxCom = context.WithValue(ctxCom, "user-agent", "")
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	cmdRoot.PersistentFlags().StringVarP(&input, "input", "i", "in.txt", "input file path, read io/stream")
	cmdRoot.PersistentFlags().StringVarP(&output, "output", "o", "out.txt", "output file path, write io/stream")
	cmdRoot.PersistentFlags().StringVarP(&env, "env", "e", "test", "env param")
	cmdRoot.PersistentFlags().StringVarP(&env, "cfg", "c", "config.yaml", "config yaml file")
}

// Execute the refresh tool
func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatalln(err)
	}
}
