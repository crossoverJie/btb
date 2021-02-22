package util

import (
	"btb/constants"
	"github.com/fatih/color"
	"regexp"
	"strings"
)

//MatchImg match images
func MatchImg(context *string) (*string, error) {
	compile, err := regexp.Compile(constants.PicPattern)
	url := compile.FindString(*context)
	return &url, err

}

func ignorePic(filePath string, line *string) bool {
	if strings.Contains(*line, constants.IgnorePic) {
		color.Red("Ignore md=[%s], pic=[%s]", filePath, *line)
		return true
	} else {
		return false
	}
}

func MatchAvailableImg(filePath string, allLine *[]string) *[]string {
	var availableImgs []string
	for _, line := range *allLine {
		if ignorePic(filePath, &line) {
			continue
		}
		url, err := MatchImg(&line)
		if err == nil && *url != "" {
			availableImgs = append(availableImgs, *url)
		}
	}
	return &availableImgs
}
