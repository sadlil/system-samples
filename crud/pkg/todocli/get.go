package todocli

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/sadlil/system-samples/crud/apis/go/crudapiv1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func newGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "get",
		Run: runGet,
	}
	return cmd
}

func runGet(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags())

	client, err := newTodoServiceClient(viper.GetString(flagTransport), viper.GetString(flagServerAddress))
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to create cleint: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	if len(args) != 1 {
		fmt.Fprintf(os.Stdout, "Exactly one todo id expected, found: %v", len(args))
		os.Exit(1)
	}

	resp, err := client.GetTodo(cmd.Root().Context(), &crudapiv1.GetTodoRequest{
		Id: args[0],
	})
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to get todo: %v", err)
		os.Exit(1)
	}

	tr := table.NewWriter()
	tr.SetOutputMirror(os.Stdout)
	tr.AppendRow(table.Row{"ID", resp.Todo.Id})
	tr.AppendRow(table.Row{"Name", resp.Todo.Name})
	tr.AppendRow(table.Row{"Description", resp.Todo.Description})
	tr.AppendRow(table.Row{"Priority", resp.Todo.Priority})
	tr.AppendRow(table.Row{"Status", resp.Todo.Status})
	tr.AppendRow(table.Row{"Created at", resp.Todo.CreatedAt.AsTime().Local().String()})
	tr.AppendRow(table.Row{"Due at", resp.Todo.CreatedAt.AsTime().Add(resp.Todo.Deadline.AsDuration()).Local().String()})

	tr.SetStyle(table.StyleColoredDark)
	tr.Render()
}
