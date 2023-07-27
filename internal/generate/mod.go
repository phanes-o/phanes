package generate

import (
	"io/ioutil"
	"log"
	"path"

	"golang.org/x/mod/modfile"
)

func getModule(pwd, project string) (string, error) {
	modContent, err := ioutil.ReadFile(path.Join(pwd, project, "go.mod"))
	if err != nil {
		log.Fatalf("Failed to read go.mod file: %v", err)
		return "", err
	}

	modFile, err := modfile.Parse("go.mod", modContent, nil)
	if err != nil {
		log.Fatalf("Failed to parse go.mod file: %v", err)
		return "", err
	}
	return modFile.Module.Mod.String(), nil
}
