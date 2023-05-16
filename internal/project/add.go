package project

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/lizhiqpxv/phanes/internal/base"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

var repoAddIgnores = []string{
	".git", ".github", "api", "README.md", "LICENSE", "go.mod", "go.sum", "third_party",
}

func (p *Project) Add(ctx context.Context, dir string, layout string, branch string, mod string) error {
	to := path.Join(dir, p.Path)

	if _, err := os.Stat(to); !os.IsNotExist(err) {
		fmt.Printf("🚫 %s already exists\n", p.Name)
		override := false
		prompt := &survey.Confirm{
			Message: "📂 Do you want to override the folder ?",
			Help:    "Delete the existing folder and create the project.",
		}
		e := survey.AskOne(prompt, &override)
		if e != nil {
			return e
		}
		if !override {
			return err
		}
		os.RemoveAll(to)
	}

	fmt.Printf("🚀 Add service %s, layout repo is %s, please wait a moment.\n\n", p.Name, layout)

	repo := base.NewRepo(layout, branch)

	if err := repo.CopyToV2(ctx, to, path.Join(mod, p.Path), repoAddIgnores, []string{path.Join(p.Path, "api"), "api"}); err != nil {
		return err
	}

	base.Tree(to, dir)

	fmt.Printf("\n🍺 Repository creation succeeded %s\n", color.GreenString(p.Name))
	fmt.Print("💻 Use the following command to add a project 👇:\n\n")

	// fmt.Println(color.WhiteString("$ cd %s", p.Name))
	// fmt.Println(color.WhiteString("$ go generate ./..."))
	// fmt.Println(color.WhiteString("$ go build -o ./bin/ ./... "))
	fmt.Println("			🤝 Thanks for using phanes")
	return nil
}
