package service

import (
	"btb/io"
	"fmt"
	"github.com/fatih/color"
	"log"
)

func DownLoadPic(markdownPath string) {

	allFile, err := io.GetAllFile(markdownPath)
	if err != nil {
		log.Fatal("read file error")
	}

	for _, filePath := range *allFile {
		//context, err := io.ReadFileAllContext(filePath)
		//if err != nil {
		//	log.Fatal(err)
		//}
		picCount := 0
		allLine, err := io.ReadFileLine(filePath)
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range *allLine {
			picture, err := io.MatchPicture(&line)
			if err == nil && *picture != "" {
				fmt.Printf("file:%v pic:%v \n", filePath, *picture)
				picCount++
			}

		}
		color.Green("Run [%v], [%v]pic successfully!!!", filePath, picCount)

	}
	color.Green("Run [%v] size successfully!!!", len(*allFile))
	if err != nil {
		log.Fatal(err)
	}
}
