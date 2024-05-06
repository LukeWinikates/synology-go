package main

import (
	"fmt"
	"github.com/LukeWinikates/synology-go/cmd/docker"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
	"os"
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
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
