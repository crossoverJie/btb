package io

import (
	"btb/constants"
	"fmt"
	"regexp"
)

//MatchPicture match pic
func MatchPicture(context string) (string, error) {
	compile, err := regexp.Compile(constants.PicPattern)
	findString := compile.FindString(context)
	fmt.Sprintf("match %s", findString)

	return findString, err

}
