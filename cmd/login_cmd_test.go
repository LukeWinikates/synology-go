package main

import (
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/LukeWinikates/synology-go/pkg/api/auth"
	"github.com/rogpeppe/go-internal/testscript"
)

var s *httptest.Server

func TestMain(m *testing.M) {
	defer s.Close()
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"synoctl": Main,
	}))
}

func Test(t *testing.T) {
	account := "fake-login"
	password := "fake-password"
	s = httptest.NewServer(auth.NewFakeAuthServerHandler(account, password))
	testscript.Run(t, testscript.Params{
		Dir: "../testdata/login",
		Setup: func(env *testscript.Env) error {
			context := []byte(strings.Join([]string{s.URL, account, password, ""}, "\n"))
			return os.WriteFile(filepath.Join(env.Getenv("WORK"), "input.txt"), context, os.ModePerm)
		},
	})
}
