package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Coingecko CLI",
	Short: "CLI for COINS",
}

func Execute() error {
	return rootCmd.Execute()
}
