package service

import (
	"btb/constants"
	"btb/util"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"log"
	"sync"
)

func DownLoadPic(markdownPath, downloadPath string) {
	wg := sync.WaitGroup{}
	allFile, err := util.GetAllFile(markdownPath)
	wg.Add(len(*allFile))

	if err != nil {
		log.Fatal("read file error")
	}

	for _, filePath := range *allFile {

		go func(filePath string) {
			allLine, err := util.ReadFileLine(filePath)
			if err != nil {
				log.Fatal(err)
			}
			availableImgs := util.MatchAvailableImg(allLine)
			bar := pb.ProgressBarTemplate(constants.PbTmpl).Start(len(*availableImgs))
			bar.Set("fileName", filePath).
				SetWidth(120)

			for _, url := range *availableImgs {
				if err != nil {
					log.Fatal(err)
				}
				err := util.DownloadFile(url, *genFullFileName(downloadPath, filePath, &url))
				if err != nil {
					log.Fatal(err)
				}
				bar.Increment()
				//fmt.Printf("Download: %v image: %v successful.\n", filePath, url)

			}
			//color.Green("Run [%v], [%v]images successfully!!!\n", filePath, len(*availableImgs))
			bar.Finish()
			wg.Done()

		}(filePath)
	}
	wg.Wait()
	color.Green("Successful handling of [%v] files.\n", len(*allFile))

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
