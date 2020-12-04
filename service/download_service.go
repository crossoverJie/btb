package service

import (
	"btb/util"
	"fmt"
	"github.com/fatih/color"
	"log"
)

func DownLoadPic(markdownPath, downloadPath string) {

	allFile, err := util.GetAllFile(markdownPath)
	if err != nil {
		log.Fatal("read file error")
	}

	for _, filePath := range *allFile {
		picCount := 0
		allLine, err := util.ReadFileLine(filePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range *allLine {
			url, err := util.MatchPicture(&line)
			if err == nil && *url != "" {
				err := util.DownloadFile(*url, *genFullFileName(downloadPath, filePath, url))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Download: %v image: %v successful.\n", filePath, *url)
				picCount++
			}

		}
		color.Green("Run [%v], [%v]images successfully!!!", filePath, picCount)

	}
	color.Green("Run [%v] size successfully!!!", len(*allFile))
	if err != nil {
		log.Fatal(err)
	}
}

// downLoadPath/filename-5cd1d11a5612f.jpg
func genFullFileName(downLoadPath, filePath string, url *string) *string {
	suffix := util.GetFileSuffix(*url)
	fullFileName := downLoadPath + "/" + util.GetFileSuffix(filePath) + "-" + suffix
	return &fullFileName
}
