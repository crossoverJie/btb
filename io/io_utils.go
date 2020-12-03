package io

import (
	"bufio"
	"io/ioutil"
	"os"
	"strings"
)

var allFile []string

func GetAllFile(filePath string) (*[]string, error) {
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
	return &allFile, err
}

func ReadFileAllContext(filePath string) (*string, error) {
	file, err := ioutil.ReadFile(filePath)
	value := string(file)
	return &value, err
}

func ReadFileLine(filePath string) (*[]string, error) {
	var allLine []string
	f, err := os.Open(filePath)
	r := bufio.NewReader(f)
	s, e := readln(r)
	for e == nil {
		allLine = append(allLine, s)
		s, e = readln(r)
	}
	return &allLine, err
}

func readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}
