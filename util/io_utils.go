package util

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
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
			GetAllFile(filePath + "/" + fi.Name())
		} else {
			if strings.HasSuffix(fi.Name(), "md") {
				allFile = append(allFile, filePath+"/"+fi.Name())
			}
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

//DownloadFile
func DownloadFile(url, fileName string) error {

	if Exist(fileName) {
		return nil
	}

	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
