package todocli

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newDoneCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "done",
		Run: runDone,
	}

	cmd.Flags().StringP("id", "i", "", "Task priority")
	cmd.MarkFlagRequired("id")
	return cmd
}

func runDone(cmd *cobra.Command, args []string) {
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
	payload.Status = crudapiv1.TodoStatus_TODO_STATUS_DONE
	_, err = client.UpdateTodo(cmd.Root().Context(), &crudapiv1.UpdateTodoRequest{
		Id:      viper.GetString("id"),
		Payload: payload,
	})
	if err != nil {
		glog.Errorf("Failed to update todo, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to update todo: %v", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Todo %v Is Done.", viper.GetString("id"))
}
