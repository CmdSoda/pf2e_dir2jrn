package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/CmdSoda/pf2e_dir2jrn/dir2jrn"
	"github.com/alexflint/go-arg"
)

func GetFiles(folder string) []string {
	var list []string = []string{}
	files, err := ioutil.ReadDir(folder)
	if err == nil {
		for _, file := range files {
			if file.IsDir() == false {
				list = append(list, file.Name())
			}
		}
	}
	return list
}

func main() {
	log.Println("pf2e_dir2jrn")

	var args struct {
		ServerPath string `arg:"positional, required"` // "Bilder/Books/Grundregelwerk/"
		LocalPath  string `arg:"positional, required"` // "S:\Data\Bilder\Books\Grundregelwerk\"
	}
	arg.MustParse(&args)

	dir2jrn.InitRand()

	folder := "D:\\Dropbox\\pathfinder\\Bilder\\Books\\Grundregelwerk\\Alles"

	files := GetFiles(folder)

	for _, file := range files {
		fmt.Println(file)
	}
}
