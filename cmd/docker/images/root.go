package images

import (
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/images"
	"github.com/spf13/cobra"
)

func Cmd(newClient func() api.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "images",
		Aliases: []string{"image"},
	}
	factory := func() images.Client { return images.NewClient(newClient()) }
	cmd.AddCommand(listCmd(factory))
	cmd.AddCommand(getCmd(factory))
	cmd.AddCommand(upgradeCmd(factory))
	cmd.AddCommand(pullCmd(factory))
	cmd.AddCommand(followUpgradeCmd(factory))
	cmd.AddCommand(followPullCmd(factory))

	return cmd
}
