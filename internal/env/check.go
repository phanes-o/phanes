package env

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func check(_ *cobra.Command, _ []string) error {
	runCheck()
	return nil
}

// Run run check env.
func runCheck() {
	deps := []string{"protoc", "protoc-gen-go", "protoc-gen-go-grpc", "protoc-gen-micro"}
	for _, d := range deps {
		_, err := LookPath(d)
		if err != nil {
			color.New(color.FgRed).Fprintln(os.Stderr, "ERROR:", err.Error())
		} else {
			fmt.Println(color.GreenString(d), "✅")
		}
	}

}
