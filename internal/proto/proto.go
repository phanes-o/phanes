package proto

import (
	"github.com/lizhiqpxv/phanes/internal/proto/add"

	"github.com/lizhiqpxv/phanes/internal/proto/client"
	"github.com/lizhiqpxv/phanes/internal/proto/server"

	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
	Run:   run,
}

func init() {
	CmdProto.AddCommand(add.CmdAdd)
	CmdProto.AddCommand(client.CmdClient)
	CmdProto.AddCommand(server.CmdServer)
}

func run(cmd *cobra.Command, args []string) {
	cmd.Help()
}
