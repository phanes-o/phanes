package main

import (
	"log"

	"github.com/phanes-o/phanes/internal/project"
	"github.com/phanes-o/phanes/internal/proto"
	"github.com/phanes-o/phanes/internal/run"
	"github.com/phanes-o/phanes/internal/upgrade"

	"github.com/spf13/cobra"
)

// release is the current kratos tool version.
const release = "v0.1.0"

var rootCmd = &cobra.Command{
	Use:     "phanes",
	Short:   "phanes: An elegant toolkit for Go microservices.",
	Long:    `phanes: An elegant toolkit for Go microservices.`,
	Version: release,
}

func init() {
	rootCmd.AddCommand(project.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
	rootCmd.AddCommand(upgrade.CmdUpgrade)
	rootCmd.AddCommand(run.CmdRun)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
