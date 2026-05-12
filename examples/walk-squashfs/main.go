package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	backendFile "github.com/diskfs/go-diskfs/backend/file"
	"github.com/diskfs/go-diskfs/filesystem/squashfs"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s path/to/filesystem.sqs", filepath.Base(os.Args[0]))
	}

	image, err := os.Open(os.Args[1])
	check(err)
	defer image.Close()

	info, err := image.Stat()
	check(err)

	sqs, err := squashfs.Read(backendFile.New(image, true), info.Size(), 0, 0)
	check(err)

	err = fs.WalkDir(sqs, ".", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			fmt.Printf("%s/\n", p)
			return nil
		}
		if d.Type().IsRegular() {
			content, err := fs.ReadFile(sqs, p)
			if err != nil {
				return err
			}
			fmt.Printf("%s:\n%s\n", p, content)
			return nil
		}
		fmt.Printf("%s [%s]\n", p, d.Type())
		return nil
	})
	check(err)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
