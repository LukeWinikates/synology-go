package main

import (
	"fmt"
	"os"

	"github.com/LukeWinikates/synology-go/cmd/docker"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(docker.Cmd(func() api.Client {
		c, err := api.NewClient(
			os.Getenv("DSM_HOST"),
			os.Getenv("DSM_ACCOUNT"),
			os.Getenv("DSM_PWD"))
		if err != nil {
			panic(err)
		}
		_, err = c.Login()
		if err != nil {
			panic(err)
		}
		return c
	}))
}

var rootCmd = &cobra.Command{
	Use: "dsmctl",
	Long: `
dsmctl is a utility for interacting with your Synology NAS from a remote terminal

`,
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
