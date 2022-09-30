package handler

import (
	"context"
	"log"

	"github.com/spf13/cobra"
)

var (
	// flag value
	ctx  context.Context
	src  string
	dest string

	env string

	// cmdRoot
	cmdRoot = &cobra.Command{
		Use:   "helpi",
		Short: "helpi is a tool as short hand helper for Dev",
		Long:  `helpi is a tool as short hand helper for developer.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			ctx = context.Background()
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	cmdRoot.PersistentFlags().StringVarP(&src, "src", "f", "src.txt", "src file path")
	cmdRoot.PersistentFlags().StringVarP(&dest, "dest", "o", "dest.txt", "dest file path")
	cmdRoot.PersistentFlags().StringVarP(&env, "env", "e", "test", "env param")
}

// Execute the refresh tool
func Execute() {
	if err := cmdRoot.Execute(); err != nil {
		log.Fatalln(err)
	}
}
