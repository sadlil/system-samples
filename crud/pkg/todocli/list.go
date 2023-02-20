package todocli

import (
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: runList,
	}

	cmd.Flags().String("priority", "", "Filter tasks by priority")
	cmd.Flags().Int("limit", 1000, "Number of todo entry to fetch")
	viper.BindPFlags(cmd.Flags())
	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	client, err := newTodoServiceClient(viper.GetString(flagTransport), viper.GetString(flagServerAddress))
	if err != nil {
		glog.Errorf("Failed to create client, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to create cleint: %v", err)
		os.Exit(1)
	}

	resp, err := client.ListTodo(cmd.Root().Context(), &crudapi.ListTodoRequest{
		Limit: int64(viper.GetInt("limit")),
	})
	if err != nil {
		glog.Errorf("Failed to list todos, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to list todos: %v", err)
		os.Exit(1)
	}

	tr := table.NewWriter()
	tr.SetOutputMirror(os.Stdout)
	tr.AppendHeader(
		table.Row{"ID", "Name", "Description", "Priority", "Created At", "Deadline"},
	)
	for _, todo := range resp.Todos {
		tr.AppendRow(
			table.Row{
				todo.Id,
				todo.Name,
				todo.Description,
				todo.Priority,
				todo.CreatedAt.AsTime().Local().String(),
				todo.CreatedAt.AsTime().Add(todo.Deadline.AsDuration()).Local().String(),
			},
		)
	}
	tr.Render()
}
