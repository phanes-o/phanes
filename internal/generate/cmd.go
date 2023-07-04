package generate

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Cmd represents the new command.
var Cmd = &cobra.Command{
	Use:   "gen",
	Short: "codeBuild all business logic code",
	Long:  "codeBuild business logic cod include bll, store, http api, grpc api, entity",
	Run:   run,
}

var (
	configName     string
	workspace      bool
	importExistMap = map[string]struct{}{}
)

func init() {
	Cmd.Flags().BoolVarP(&workspace, "workspace", "w", false, "enable workspace")
	Cmd.Flags().StringVarP(&configName, "config", "c", "generator.go", "codeBuild config file")
}

func run(cmd *cobra.Command, args []string) {
	var (
		err       error
		pwd       string
		generator *Generator
	)
	// Check args

	// If enable workspace. check workspace environment. means Phanes command must run in go workspace
	if workspace {
		if !checkEnvironment() {
			fmt.Println(color.RedString(fmt.Sprintf("ERROR: can not detection go workspace. please run `go work init` to init go workspace")))
			return
		}
	}

	if pwd, err = os.Getwd(); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: %s", err)))
		os.Exit(1)
	}

	if !fileExists(path.Join(pwd, configName)) {
		fmt.Println(color.RedString("Generate config file does exist!"))
		os.Exit(1)
	}

	// Build Generator
	// Parse template and replace it
	if generator, err = ReadSource(configName); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: read generator.go error %s", err)))
		return
	}
	// Generate nad write code to destination file.
	if err = generator.Generate(); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: Generate failed %s", err)))
		return
	}
}
