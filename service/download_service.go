package service

import (
	"btb/util"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"log"
	"sync"
)

func DownLoadPic(markdownPath, downloadPath string) {
	wg := sync.WaitGroup{}
	allFile, err := util.GetAllFile(markdownPath)
	wg.Add(len(*allFile))
	bar := pb.StartNew(len(*allFile))
	if err != nil {
		log.Fatal("read file error")
	}

	for _, filePath := range *allFile {

		go func(filePath string) {
			picCount := 0
			allLine, err := util.ReadFileLine(filePath)
			if err != nil {
				log.Fatal(err)
			}
			availableImgs := util.MatchAvailableImg(allLine)
			for _, url := range *availableImgs {
				if err != nil {
					log.Fatal(err)
				}
				err := util.DownloadFile(url, *genFullFileName(downloadPath, filePath, &url))
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Download: %v image: %v successful.\n", filePath, url)
				picCount++

			}
			color.Green("Run [%v], [%v]images successfully!!!", filePath, len(*availableImgs))
			wg.Done()
			bar.Increment()
		}(filePath)
	}
	wg.Wait()
	bar.Finish()
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
