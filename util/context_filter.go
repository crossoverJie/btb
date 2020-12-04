package util

import (
	"btb/constants"
	"regexp"
)

//MatchPicture match pic
func MatchPicture(context *string) (*string, error) {
	compile, err := regexp.Compile(constants.PicPattern)
	url := compile.FindString(*context)
	return &url, err

}
