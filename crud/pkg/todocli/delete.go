package todocli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

func newDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		Run: runDelete,
	}
	return cmd
}

func runDelete(cmd *cobra.Command, args []string) {
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

	_, err = client.DeleteTodo(cmd.Root().Context(), &crudapi.DeleteTodoRequest{
		Id: args[0],
	})
	if err != nil {
		fmt.Fprintf(os.Stdout, "Failed to delete todo: %v", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Todo %v deleted", args[0])
}
