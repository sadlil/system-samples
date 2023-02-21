package todocli

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/protobuf/types/known/durationpb"
	"sadlil.com/samples/crud/apis/go/crudapi"
)

func newCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create",
		Run: runCreate,
	}

	cmd.Flags().StringP("name", "n", "", "Task priority")
	cmd.Flags().String("desc", "", "Task description")
	cmd.Flags().StringP("priority", "p", "P1", "Task priority")
	cmd.Flags().DurationP("deadline", "d", time.Hour, "Task finish deadline")

	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("deadline")
	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {
	viper.BindPFlags(cmd.Flags())

	client, err := newTodoServiceClient(viper.GetString(flagTransport), viper.GetString(flagServerAddress))
	if err != nil {
		glog.Errorf("Failed to create client, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to create cleint: %v", err)
		os.Exit(1)
	}
	defer client.Close()

	_, err = client.CreateTodo(cmd.Root().Context(), &crudapi.CreateTodoRequest{
		Todo: &crudapi.Todo{
			Name:        viper.GetString("name"),
			Description: viper.GetString("desc"),
			Priority:    viper.GetString("priority"),
			Deadline:    durationpb.New(viper.GetDuration("deadline")),
		},
	})
	if err != nil {
		glog.Errorf("Failed to create todo, reason: %v", err)
		fmt.Fprintf(os.Stdout, "Failed to create todo: %v", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Todo %v created.", viper.GetString("name"))
}
