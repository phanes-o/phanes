package register

import (
	"fmt"
	"os"
	"path"

	"github.com/fatih/color"
	"github.com/phanes-o/phanes/internal/utils"
	"github.com/spf13/cobra"
)

// Cmd represents the new command.
var Cmd = &cobra.Command{
	Use:   "register",
	Short: "codeBuild all business logic code",
	Long:  "codeBuild business logic cod include bll, store, http api, grpc api, entity",
	Run:   run,
}

var (
	registryType string
	project      string
)

func init() {
	Cmd.Flags().StringVarP(&registryType, "registry_type", "t", "etcd", "registry type")
	Cmd.Flags().StringVarP(&project, "project", "p", "", "your project name")
}

func run(cmd *cobra.Command, args []string) {
	var (
		pwd   string
		err   error
		bytes []byte
	)
	// check env
	if project == "" {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: please spacify project name")))
		return
	}

	if pwd, err = os.Getwd(); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: %s", err)))
		os.Exit(1)
	}

	if !utils.FileExists(path.Join(pwd, project)) {
		fmt.Println(color.RedString("project does not exist!"))
		os.Exit(1)
	}

	// parse package and ast file
	if bytes, err = parsePackage(pwd, project); err != nil {
		fmt.Println(color.RedString("parse file error"))
		os.Exit(1)
	}

	fileName := path.Join(pwd, project, "assistant", "resource_register.go")

	// save code
	if !utils.FileExists(fileName) {
		if err = utils.WriteFile(fileName, bytes); err != nil {
			fmt.Println(color.RedString("write code to file error"))
			os.Exit(1)
		}
	} else {
		fmt.Println(color.RedString("register file already exist!"))
	}
}
