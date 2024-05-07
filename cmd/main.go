package main

import (
	"fmt"
	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
	"log"
	"os"

	"github.com/LukeWinikates/synology-go/cmd/docker"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
)

func rootCmd(c *config) *cobra.Command {
	cmd := &cobra.Command{
		Use: "dsmctl",
		Long: `
dsmctl is a utility for interacting with your Synology NAS from a remote terminal

`}
	cmd.AddCommand(docker.Cmd(func() api.Client {
		apiClient, err := api.NewClient(c.Host, c.Account, c.Password)
		if err != nil {
			panic(err)
		}
		_, err = apiClient.Login()
		if err != nil {
			panic(err)
		}
		return apiClient
	}))
	return cmd
}

type config struct {
	Host     string `yaml:"host"`
	Account  string `yaml:"account"`
	Password string `yaml:"password"`
}

func main() {
	configFilePath, err := xdg.SearchConfigFile(".dsmctl.yaml")
	if err != nil {
		log.Println(`
couldn't find a .dsmctl.yaml file. You'll need to create one with the following format:
host: "http://$MY_SYNOLOGY_HOST:5000"
account: "my-user"
password: "password"
`)
		log.Fatal(err)
	}
	open, err := os.Open(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
	var c *config
	yaml.NewDecoder(open).Decode(&c)
	if err = rootCmd(c).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
