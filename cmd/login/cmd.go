package login

import (
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"gopkg.in/yaml.v3"
)

type SessionProvider struct {
	Host            string `yaml:"host"`
	SessionID       string `yaml:"sessionID"`
	sessionFilePath string
}

func (sp *SessionProvider) Login(acct, pwd string) error {
	c, err := api.NewClient(sp.Host)
	if err != nil {
		return err
	}
	response, err := c.Login(acct, pwd)
	if err != nil {
		return err
	}
	sp.SessionID = response.Data.Sid
	return nil
}

func (sp *SessionProvider) SaveSession() error {
	sessionFile, err := os.Create(sp.sessionFilePath)
	if err != nil {
		return err
	}
	defer sessionFile.Close()
	return yaml.NewEncoder(sessionFile).Encode(&sp)
}

func Cmd(sp *SessionProvider) *cobra.Command {
	return &cobra.Command{
		Use: "login",
		RunE: func(_ *cobra.Command, _ []string) error {
			var host, acct string
			if sp.Host == "" {
				fmt.Println("enter your Synology/DSM account host:")
				if _, err := fmt.Scanln(&host); err != nil {
					return err
				}
				sp.Host = host
			}
			fmt.Println("enter your Synology/DSM account name:")
			if _, err := fmt.Scanln(&acct); err != nil {
				return err
			}
			fmt.Println("password:")
			pwdBytes, err := term.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return fmt.Errorf("no password entered")
			}
			err = sp.Login(acct, string(pwdBytes))
			if err != nil {
				return err
			}
			err = sp.SaveSession()
			if err != nil {
				fmt.Fprintf(os.Stderr, "login failed: %s\n", err.Error())
			} else {
				fmt.Println("login succeeded")
			}
			return err
		},
	}
}

func NewSessionProvider(sessionFilePath string) *SessionProvider {
	var sp *SessionProvider
	file, err := os.Open(sessionFilePath)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.NewDecoder(file).Decode(&sp)
	if err != nil {
		sp = &SessionProvider{}
	}
	sp.sessionFilePath = sessionFilePath
	return sp
}
