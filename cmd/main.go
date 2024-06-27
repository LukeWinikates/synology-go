package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/LukeWinikates/synology-go/cmd/docker"
	"github.com/LukeWinikates/synology-go/cmd/login"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/api/auth"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func ensureConfigFile() (string, error) {
	configFilePath, err := xdg.SearchConfigFile(".synoctl.yaml")
	if err == nil {
		return configFilePath, nil
	}
	filePath := filepath.Join(xdg.ConfigHome, ".synoctl.yaml")
	log.Printf("Couldn't find a .synoctl.yaml file. Creating an empty one at %s", filePath)
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	return filePath, file.Close()
}

func rootCmd(sp *login.SessionProvider) *cobra.Command {
	debug := false

	cmd := &cobra.Command{
		Use: "synoctl",
		Long: `
synoctl is a utility for interacting with your Synology NAS from a remote terminal
`,
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			if debug {
				development, err := zap.NewDevelopment()
				api.SetLogger(development)
				return err
			}
			return nil
		},
	}
	cmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "")

	cmd.AddCommand(docker.Cmd(newAPIClient(sp)))
	cmd.AddCommand(login.Cmd(sp))
	return cmd
}

func newAPIClient(sp *login.SessionProvider) func() api.Client {
	return func() api.Client {
		return api.NewClient(sp.Host, auth.NewSessionAuthorizer(sp.SessionID))
	}
}

func main() {
	file, err := ensureConfigFile()
	if err != nil {
		log.Fatal(err)
	}

	if err = rootCmd(login.NewSessionProvider(file)).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
