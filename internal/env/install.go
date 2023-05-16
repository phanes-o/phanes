package env

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/lizhiqpxv/phanes/internal/base"
	"github.com/spf13/cobra"
)

func install(_ *cobra.Command, _ []string) error {
	runInstall()
	return nil
}

func runInstall() {
	err := base.GoInstall(
		"go-micro.dev/v4/cmd/protoc-gen-micro@v4",
		"google.golang.org/protobuf/cmd/protoc-gen-go@latest",
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest",
	)
	if err != nil {
		fmt.Println(color.RedString("install error: %s", err.Error()))
	}
}
