package util

import "strings"

func GetFileSuffix(fileName string) string {
	index := strings.LastIndex(fileName, "/")
	fileName = fileName[index+1:]
	return fileName
}
