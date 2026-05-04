package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tmz",
	Short: "tmz is a CLI tool to calculate time & times across different timezones",
	Long:  `A fast and flexible timezone conversion tool built with Cobra.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
