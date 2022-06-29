package upgrade

import (
	"fmt"

	"github.com/phanes-o/phanes/cmd/phanes/internal/base"

	"github.com/spf13/cobra"
)

// CmdUpgrade represents the upgrade command.
var CmdUpgrade = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade the phanes tools",
	Long:  "Upgrade the phanes tools. Example: phanes upgrade",
	Run:   Run,
}

// Run upgrade the kratos tools.
func Run(cmd *cobra.Command, args []string) {
	err := base.GoInstall(
		"go-micro.dev/v4/cmd/protoc-gen-micro@v4",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	)
	if err != nil {
		fmt.Println(err)
	}
}
