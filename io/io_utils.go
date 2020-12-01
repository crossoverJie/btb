package io

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var allFile []string

func GetAllFile(filePath string) ([]string, error) {
	dir, err := ioutil.ReadDir(filePath)
	for _, fi := range dir {
		if strings.HasPrefix(fi.Name(), ".") {
			continue
		}
		if fi.IsDir() {
			GetAllFile(filePath + "/" + fi.Name() + "/")
		} else {
			allFile = append(allFile, filePath+"/"+fi.Name())
		}
	}
	return allFile, err
}

func ReadFileContext(filePath string) (string, error) {
	file, openErr := os.Open(filePath)
	all, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return string(all), openErr

}
