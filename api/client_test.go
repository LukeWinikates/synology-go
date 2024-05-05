package api

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	client, _ := NewClient(
		os.Getenv("DSM_HOST"),
		os.Getenv("DSM_ACCOUNT"),
		os.Getenv("DSM_PWD"))
	_, err := client.Login()
	assert.NoError(t, err)
	info, err := client.GetInfo()
	assert.NoError(t, err)
	fmt.Println(info.Data)
}
