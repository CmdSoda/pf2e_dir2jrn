package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/CmdSoda/pf2e_dir2jrn/dir2jrn"
	"github.com/alexflint/go-arg"
)

var frontTemplate string = "{\"name\": \"$BOOKNAME$\", \"pages\": ["
var pageTemplate string = "{\"sort\": $PAGESORT$,\"name\": \"$PAGETITLE$\",\"type\": \"image\",\"_id\": \"$PAGEID$\",\"title\": {\"show\": true,\"level\": 1},\"image\": {\"caption\": \"\"},\"text\": {\"format\": 1},\"video\": {\"controls\": true,\"volume\": 0.5},\"src\":\"$PAGESERVERPATHFILENAME$\",\"system\": {},\"ownership\": {\"default\": -1},\"flags\": {}}"
var endTemplate string = "], \"flags\": {    \"exportSource\": {      \"world\": \"$WORLD$\",      \"system\": \"pf2e\",      \"coreVersion\": \"10.290\",     \"systemVersion\": \"4.3.4\"    }  },  \"_stats\": {    \"systemId\": \"pf2e\",    \"systemVersion\": \"4.3.4\",    \"coreVersion\":\"10.290\",    \"createdTime\": 1668237338236,    \"modifiedTime\": 1668300285243,    \"lastModifiedBy\":\"EoVp4BXD44LN4CZA\"  }}"

func GetFiles(folder string) ([]string, error) {
	var list []string = []string{}
	files, err := ioutil.ReadDir(folder)
	if err == nil {
		for _, file := range files {
			if file.IsDir() == false {
				list = append(list, file.Name())
			}
		}
	} else {
		return []string{}, err
	}
	return list, nil
}

func createServerPathFile(serverpath string, file string) string {
	return serverpath + file
}

func SortFileNameAscend(files []string) {
	sort.Slice(files, func(i, j int) bool {
		a := files[i][:strings.Index(files[i], ".")]
		an, _ := strconv.Atoi(a)
		b := files[j][:strings.Index(files[j], ".")]
		bn, _ := strconv.Atoi(b)
		return an < bn
	})
}

func GetFilenameWithoutExt(fullname string) string {
	return fullname[:strings.Index(fullname, ".")]
}

func readIndexFile(indexfilename string) map[int]string {
	var index map[int]string = map[int]string{}

	file, err := os.Open(indexfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ":")
		ln, err := strconv.Atoi(split[0])
		if err == nil {
			index[ln] = split[1]
		}
	}

	return index
}

func main() {
	log.Println("pf2e_dir2jrn")

	var args struct {
		BookName   string `arg:"positional, required"` // "GRW"
		World      string `arg:"positional, required"` // "story1"
		ServerPath string `arg:"positional, required"` // "Bilder/Books/Grundregelwerk/"
		LocalPath  string `arg:"positional, required"` // "S:\Data\Bilder\Books\Grundregelwerk\"
		IndexFile  string `arg:"-i"`                   // "-i grw.txt"
	}
	arg.MustParse(&args)

	dir2jrn.InitRand()

	files, err := GetFiles(args.LocalPath)
	if err != nil {
		log.Fatalf("LocalPath problem: %s\n", args.LocalPath)
	}

	//fmt.Println(files)

	SortFileNameAscend(files)

	var book strings.Builder

	front := frontTemplate
	front = strings.ReplaceAll(front, "$BOOKNAME$", args.BookName)
	book.WriteString(front)

	var index map[int]string = map[int]string{}
	if args.IndexFile != "" {
		index = readIndexFile(args.IndexFile)
	}
	//fmt.Println(index)

	sort := 100000
	first := true
	for _, file := range files {
		if !first {
			book.WriteString(",")
		}
		first = false
		entry := pageTemplate
		entry = strings.ReplaceAll(entry, "$PAGESORT$", strconv.Itoa(sort))
		filestr := GetFilenameWithoutExt(file)
		//fmt.Println(filestr)
		nr, _ := strconv.Atoi(filestr)
		//fmt.Println(nr)
		idx, ok := index[nr]
		if ok {
			entry = strings.ReplaceAll(entry, "$PAGETITLE$", idx)
		} else {
			entry = strings.ReplaceAll(entry, "$PAGETITLE$", filestr)
		}
		entry = strings.ReplaceAll(entry, "$PAGEID$", dir2jrn.CreateID())
		entry = strings.ReplaceAll(entry, "$PAGESERVERPATHFILENAME$", createServerPathFile(args.ServerPath, file))
		book.WriteString(entry)
		sort += 100000
	}

	end := endTemplate
	end = strings.ReplaceAll(end, "$WORLD$", args.World)
	book.WriteString(end)

	fmt.Println(book.String())
}
