package upgrade

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/lizhiqpxv/phanes/internal/base"
	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the phanes tools",
	Long:  "Upgrade the phanes tools. Example: phanes upgrade",
	Run:   Run,
}

// Run upgrade the phanes tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"go-micro.dev/v4/cmd/protoc-gen-micro@v4",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	)
	if err != nil {
		fmt.Println(color.RedString("install error: %s", err.Error()))

	}
}
