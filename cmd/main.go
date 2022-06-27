package main

import (
	"log"
	"phanes/internal/project"
	"phanes/internal/proto"
	"phanes/internal/run"
	"phanes/internal/upgrade"

	"github.com/spf13/cobra"
)

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
