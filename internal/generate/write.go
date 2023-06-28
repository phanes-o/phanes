package generate

import "os"

func writeFile(filePath string, data []byte) error {
	var (
		f   *os.File
		err error
	)
	if f, err = os.Create(filePath); err != nil {
		return err
	}
	defer f.Close()

	if _, err = f.Write(data); err != nil {
		return err
	}
	return nil
}
