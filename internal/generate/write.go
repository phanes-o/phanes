package generate

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func writeFile(filePath string, data []byte) error {
	var (
		f   *os.File
		err error
	)

	dirPath := filepath.Dir(filePath)

	if err = os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: Can not create dir [%s] code", dirPath)), "❌  ")
		return err
	}

	if f, err = os.Create(filePath); err != nil {
		fmt.Println(color.RedString(fmt.Sprintf("ERROR: Can not create file [%s] code", dirPath)), "❌  ")
		return err
	}
	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}
