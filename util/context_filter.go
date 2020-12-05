package util

import (
	"btb/constants"
	"regexp"
)

//MatchImg match images
func MatchImg(context *string) (*string, error) {
	compile, err := regexp.Compile(constants.PicPattern)
	url := compile.FindString(*context)
	return &url, err

}

func MatchAvailableImg(allLine *[]string) *[]string {
	var availableImgs []string
	for _, line := range *allLine {
		url, err := MatchImg(&line)
		if err == nil && *url != "" {
			availableImgs = append(availableImgs, *url)
		}
	}
	return &availableImgs
}
