package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Panic(err)
	}

	var folderName = flag.String("dir", "", "The folder that will be deleted")

	flag.Parse()

	if *folderName == "" {
		log.Panic("You must specify a folder name")
	}

	log.Printf("Current working directory: %s", cwd)
	log.Printf("Folder to be deleted: %s", *folderName)

	deleteList := make([]string, 0)

	walk(cwd, *folderName, nil, &deleteList)

	for _, file := range deleteList {
		log.Printf("Deleting file: %s", file)

		err := os.RemoveAll(file)
		if err != nil {
			log.Printf("Error deleting file: %s", err)
		}
	}

	log.Printf("Done")
}

func walk(base, target string, entry fs.DirEntry, output *[]string) error {
	var fullPath string
	if entry == nil {
		fullPath = base
	} else {
		fullPath = path.Join(base, entry.Name())
	}

	if entry != nil && entry.IsDir() && entry.Name() == target {
		*output = append(*output, fullPath)

		return nil
	}

	if entry == nil || entry.IsDir() {
		dirEntry, err := os.ReadDir(fullPath)
		if err != nil {
			return nil
		}

		for _, dir := range dirEntry {
			walk(fullPath, target, dir, output)
		}
	}

	return nil
}
