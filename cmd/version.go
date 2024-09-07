package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func versionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the current version and exit",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("synoctl %s, commit %s, built at %s", version, commit, date)
		}}
}
