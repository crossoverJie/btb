package io

import (
	"btb/constants"
	"regexp"
)

//MatchPicture match pic
func MatchPicture(context *string) (*string, error) {
	compile, err := regexp.Compile(constants.PicPattern)
	findString := compile.FindString(*context)
	return &findString, err

}
