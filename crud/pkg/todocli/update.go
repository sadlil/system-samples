package todocli

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/durationpb"
)

func newUpdateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update",
		Run: runUpdate,
	}

	cmd.Flags().StringP("id", "i", "", "Task priority")
	cmd.Flags().StringP("name", "n", "", "Task priority")
	cmd.Flags().String("desc", "", "Task description")
	cmd.Flags().StringP("priority", "p", "P1", "Task priority")
	cmd.Flags().DurationP("deadline", "d", time.Hour, "Task finish deadline")

	cmd.MarkFlagRequired("id")
	return cmd
}

func runUpdate(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags())

	client, err := newTodoServiceClient(viper.GetString(flagTransport), viper.GetString(flagServerAddress))
	if err != nil {
		glog.Errorf("Failed to create client, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to create cleint: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	resp, err := client.GetTodo(cmd.Root().Context(), &crudapiv1.GetTodoRequest{
		Id: viper.GetString("id"),
	})
	if err != nil {
		glog.Errorf("Failed to get todo, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to get todo: %v", err)
		os.Exit(1)
	}

	payload := resp.Todo
	if name := viper.GetString("name"); len(name) > 0 {
		payload.Name = name
	}

	if desc := viper.GetString("desc"); len(desc) > 0 {
		payload.Description = desc
	}

	if p := viper.GetString("priority"); len(p) > 0 {
		payload.Priority = p
	}

	if dl := viper.GetDuration("deadline"); dl > 0 {
		payload.Deadline = durationpb.New(dl)
	}

	_, err = client.UpdateTodo(cmd.Root().Context(), &crudapiv1.UpdateTodoRequest{
		Id:      viper.GetString("id"),
		Payload: payload,
	})
	if err != nil {
		glog.Errorf("Failed to update todo, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to update todo: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Todo %v updated.", viper.GetString("id"))
}
