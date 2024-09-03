package images

import (
	"strings"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/internal/docker"
	"github.com/LukeWinikates/synology-go/pkg/docker/images"
)

func imagePrinter() internal.TableWriter[*images.Image] {
	return internal.NewTableWriter[*images.Image]([]string{
		"ID",
		"Name",
		"Tag",
	}, func(item *images.Image) []string {
		return []string{item.ID, item.Image, item.Tag}
	})
}

func listImagePrinter() internal.TableWriter[*images.ListImage] {
	return internal.NewTableWriter[*images.ListImage]([]string{
		"ID",
		"Repository (short)",
		"Tags",
		"Upgradeable",
	}, func(item *images.ListImage) []string {
		upgradable := ""
		if item.Upgradable {
			upgradable = "*"
		}
		return []string{
			item.ID,
			docker.RepositoryShortName(item.Repository),
			strings.Join(item.Tags, ", "),
			upgradable,
		}
	})
}
