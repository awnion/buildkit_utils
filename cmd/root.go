package cmd

import (
	"buildkit_utils/cmd/hash"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "buildkit-gosh",
	Short:         "DO NOT EXECUTE THIS BINARY MANUALLY",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func Execute() {
	rootCmd.AddCommand(hash.HashCmd())
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
