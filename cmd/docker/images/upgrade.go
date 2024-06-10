package images

import (
	"fmt"
	"time"

	"github.com/LukeWinikates/synology-go/internal"
	"github.com/LukeWinikates/synology-go/pkg/api"
	"github.com/LukeWinikates/synology-go/pkg/docker/images"
	"github.com/spf13/cobra"
)

func upgradeCmd(clientFactory func() images.Client) *cobra.Command {
	var repository string
	cmd := &cobra.Command{
		Use: "upgrade",
		RunE: func(_ *cobra.Command, _ []string) error {
			client := clientFactory()
			taskResponse, err := client.StartUpgradeCheck(repository)
			if err != nil {
				return err
			}
			fmt.Printf("image upgrade task: %s\n", taskResponse.Data.TaskID)
			return pollUpgradeTask(client, taskResponse.Data.TaskID)
		},
	}
	cmd.Flags().StringVarP(&repository, "repository", "r", "", "image repository")
	internal.Must(cmd.MarkFlagRequired("repository"))
	return cmd
}

func pullCmd(clientFactory func() images.Client) *cobra.Command {
	var repository, tag string
	cmd := &cobra.Command{
		Use: "pull",
		RunE: func(_ *cobra.Command, _ []string) error {
			client := clientFactory()
			taskResponse, err := client.StartPull(repository, tag)
			if err != nil {
				return err
			}
			return pollPullTask(client, taskResponse.Data.TaskID)
		},
	}
	cmd.Flags().StringVarP(&repository, "repository", "r", "", "image repository")
	cmd.Flags().StringVarP(&tag, "tag", "t", "", "image tag")
	internal.Must(cmd.MarkFlagRequired("repository"))
	internal.Must(cmd.MarkFlagRequired("tag"))
	return cmd
}

func followUpgradeCmd(clientFactory func() images.Client) *cobra.Command {
	var taskID string
	cmd := &cobra.Command{
		Use: "follow-upgrade",
		RunE: func(_ *cobra.Command, _ []string) error {
			return pollUpgradeTask(clientFactory(), taskID)
		},
	}
	cmd.Flags().StringVarP(&taskID, "task", "t", "", "task id")
	internal.Must(cmd.MarkFlagRequired("task"))
	return cmd
}

func pollUpgradeTask(client images.Client, taskID string) error {
	fmt.Printf("image upgrade task: %s\n", taskID)
	fmt.Println("polling image upgrade task status")
	return poll(
		3*time.Second,
		func() (*api.ResponseWrapper[*images.UpgradeStatus], error) {
			return client.GetUpgradeTaskStatus(taskID)
		}, func(response *images.UpgradeStatus) bool {
			return response.Finished
		})
}

func followPullCmd(clientFactory func() images.Client) *cobra.Command {
	var taskID string
	cmd := &cobra.Command{
		Use: "follow-pull",
		RunE: func(_ *cobra.Command, _ []string) error {
			return pollPullTask(clientFactory(), taskID)
		},
	}
	cmd.Flags().StringVarP(&taskID, "task", "t", "", "task id")
	internal.Must(cmd.MarkFlagRequired("task"))
	return cmd
}

func pollPullTask(client images.Client, taskID string) error {
	fmt.Printf("image pull task: %s\n", taskID)
	fmt.Println("polling image pull task status")
	return poll(
		3*time.Second,
		func() (*api.ResponseWrapper[*images.PullStatus], error) {
			return client.GetPullStatus(taskID)
		}, func(response *images.PullStatus) bool {
			return response.Finished
		})
}

func poll[T any](duration time.Duration, apiCall func() (*api.ResponseWrapper[T], error), finishedPredicate func(response T) bool) error {
	fmt.Printf("every: %s\n", duration.String())
	for {
		resp, err := apiCall()
		if err != nil {
			return err
		}
		if finishedPredicate(resp.Data) {
			fmt.Println("done")
			break
		}
		time.Sleep(duration)
	}
	return nil
}
