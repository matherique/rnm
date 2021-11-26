package main

import (
	"flag"
	"log"

	"github.com/matherique/rnm/internal/rename"
	"github.com/matherique/rnm/pkg/files"
)

func main() {
	name := flag.String("n", "", "Name")
	pref := flag.String("p", "", "Prefix to add")
	suf := flag.String("s", "", "Suffix to add")
	dv := flag.String("d", "", "Devider")

	folder_path := flag.String("f", ".", "Folder path")

	flag.Parse()

	if *name == "" {
		log.Fatal("Name is required")
	}

	if *pref == "" && *suf == "" {
		log.Fatal("Prefix or suffix is required")
	}

	// list all files from a folder
	files, err := files.List(*folder_path)

	if err != nil {
		log.Fatal(err)
	}

	// rename files
	if err := rename.Do(files, *name, *pref, *suf, *dv); err != nil {
		log.Fatal(err)
	}
}
