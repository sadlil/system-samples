package todocli

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	flagTransport     = "transport"
	flagServerAddress = "addr"
)

func init() {
	flag.Set("stderrthreshold", "FATAL")
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "todocli",
		Example: "todocli get <id>",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			glog.Infof("todocli command started, args %v", args)
			cmd.Help()
		},
	}

	cmd.PersistentFlags().StringP(flagTransport, "t", "http", "Transport to use for the coneection to server, http or grpc")
	cmd.PersistentFlags().StringP(flagServerAddress, "a", "http://localhost:6002", "Server address")
	viper.BindPFlags(cmd.PersistentFlags())

	cmd.AddCommand(newListCmd())

	// parse flags beyond subcommand - get around go flag 'limitations':
	// "Flag parsing stops just before the first non-flag argument" (ref: https://pkg.go.dev/flag#hdr-Command_line_flag_syntax)
	pflag.CommandLine.ParseErrorsWhitelist.UnknownFlags = true
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	// avoid 'pflag: help requested' error, as help will be defined later by cobra cmd.Execute()
	pflag.BoolP("help", "h", false, "")
	pflag.Parse()

	return cmd
}
