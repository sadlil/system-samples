package todocli

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		Run: runList,
	}

	// cmd.Flags().String("priority", "", "Filter tasks by priority")
	cmd.Flags().Int("limit", 1000, "Number of todo entry to fetch")

	return cmd
}

func runList(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags())

	client, err := newTodoServiceClient(viper.GetString(flagTransport), viper.GetString(flagServerAddress))
	if err != nil {
		glog.Errorf("Failed to create client, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to create cleint: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	resp, err := client.ListTodo(cmd.Root().Context(), &crudapiv1.ListTodoRequest{
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
		table.Row{"ID", "Name", "Description", "Priority", "Created At", "Deadline", "Status", "Due In"},
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
				todo.Status.Enum().String(),
				time.Until(todo.CreatedAt.AsTime().Add(todo.Deadline.AsDuration())).String(),
			},
		)
	}
	tr.SetColumnConfigs([]table.ColumnConfig{
		{
			Name:  "ID",
			Align: text.AlignCenter,
		},
		{
			Name:  "Priority",
			Align: text.AlignCenter,
		},
	})
	tr.SetStyle(table.StyleColoredDark)
	tr.Render()
}
