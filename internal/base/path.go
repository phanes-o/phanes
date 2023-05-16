package base

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/lizhiqpxv/phanes/internal/global"
)

var (
	ignoresname = []byte("github.com/lizhiqpxv")
)

func Home() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	home := path.Join(dir, ".phanes")
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func HomeWithDir(dir string) string {
	home := path.Join(Home(), dir)
	if _, err := os.Stat(home); os.IsNotExist(err) {
		if err := os.MkdirAll(home, 0o700); err != nil {
			log.Fatal(err)
		}
	}
	return home
}

func readfileForLine(file io.Reader, replaces []string) []byte {
	buf := bytes.Buffer{}
	fileScanner := bufio.NewScanner(file)
	// read line by line
	for fileScanner.Scan() {
		line := fileScanner.Bytes()
		if bytes.Contains(line, ignoresname) {
			buf.Write(line)
			buf.WriteString("\n")
		} else {
			var old string
			for i, next := range replaces {
				if i%2 == 0 {
					old = next
					continue
				}
				bufs := bytes.ReplaceAll(line, []byte(old), []byte(next))
				buf.Write(bufs)
				buf.WriteString("\n")
			}
		}

	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	return buf.Bytes()

}

func copyFile(src, dst string, replaces []string) error {
	var err error
	srcinfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	fs, err := os.Open(src)
	if err != nil {
		return err
	}
	defer fs.Close()
	buf := readfileForLine(fs, replaces)
	return os.WriteFile(dst, buf, srcinfo.Mode())
}

func copyDir(src, dst string, replaces, ignores []string) error {
	var err error
	var fds []os.DirEntry
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = os.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		if hasSets(fd.Name(), ignores) {
			continue
		}
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())
		var e error
		if fd.IsDir() {
			e = copyDir(srcfp, dstfp, replaces, ignores)
		} else {
			e = copyFile(srcfp, dstfp, replaces)
		}
		if e != nil {
			return e
		}
	}
	return nil
}

func hasSets(name string, sets []string) bool {
	for _, ig := range sets {
		if ig == name {
			return true
		}
	}
	return false
}

func Tree(path string, dir string) {
	_ = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && info != nil && !info.IsDir() {
			if global.VerboseOut {
				fmt.Printf("%s %s (%v bytes)\n", color.GreenString("CREATED"), strings.Replace(path, dir+"/", "", -1), info.Size())
			}

		}
		return nil
	})
}
