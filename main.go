package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var format = flag.String("f", "csv", "output format: \"csv\", \"tsv\" or \"json\"")
	flag.Parse()
	rootDir := "."
	if len(flag.Args()) != 0 {
		rootDir = flag.Arg(0)
	}
	if *format != "csv" && *format != "tsv" && *format != "json" {
		log.Fatal("-f option must be \"csv\", \"tsv\" or \"json\"")
		os.Exit(1)
	}
	results := listFiles(rootDir)
	total := CountResult{FilePath: "total", FileType: "", Code: 0, Empty: 0, Comment: 0}
	for _, result := range results {
		total.Comment = total.Comment + result.Comment
		total.Empty = total.Empty + result.Empty
		total.Code = total.Code + result.Code
		printCountResult(result, *format)
	}
	printCountResult(total, *format)
}

func printCountResult(result CountResult, format string) {
	switch format {
	case "csv":
		fmt.Println(fmt.Sprintf("%s,%s,%d,%d,%d", result.FilePath, result.FileType, result.Code, result.Empty, result.Comment))
	case "tsv":
		fmt.Println(fmt.Sprintf("%s\t%s\t%d\t%d\t%d", result.FilePath, result.FileType, result.Code, result.Empty, result.Comment))
	case "json":
		json, err := json.Marshal(result)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		fmt.Println(string(json))
	}
}

func listFiles(path string) []CountResult {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	results := []CountResult{}
	for _, file := range files {
		if !strings.HasPrefix(file.Name(), ".") {
			filePath := filepath.Join(path, file.Name())
			if file.IsDir() {
				results = append(results, listFiles(filePath)...)
			} else {
				counter := getCounter(filePath)
				results = append(results, counter.Count(filePath))
			}
		}
	}
	return results
}
