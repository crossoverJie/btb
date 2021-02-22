package service

import (
	"btb/constants"
	"btb/util"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"log"
	"sync"
)

func DownLoadPic(markdownPath, downloadPath string, uploadService UploadService) *map[string]map[string]string {
	wg := sync.WaitGroup{}
	allFile, err := util.GetAllFile(markdownPath)
	wg.Add(len(*allFile))

	if err != nil {
		log.Fatal("read file error")
	}

	var picMapping = make(map[string]map[string]string)

	for _, filePath := range *allFile {

		go func(filePath string) {
			allLine, err := util.ReadFileLine(filePath)
			if err != nil {
				log.Fatal(err)
			}
			availableImgs := util.MatchAvailableImg(filePath, allLine)
			bar := pb.ProgressBarTemplate(constants.PbTmpl).Start(len(*availableImgs))
			bar.Set("fileName", filePath).
				SetWidth(120)

			var picMap = make(map[string]string)

			for _, url := range *availableImgs {
				if err != nil {
					log.Fatal(err)
				}
				filePathName := *genFullFileName(downloadPath, filePath, &url)
				err := util.DownloadFile(url, filePathName)
				if err != nil {
					log.Fatal(err)
				}

				//Upload file
				if constants.AppModel == constants.Replace {
					uploadPath, err := uploadService.upload(filePathName)
					if err != nil {
						log.Fatalf("Upload file [%s] error %v", filePathName, err)
						return
					}
					// Cache url
					picMap[url] = uploadPath
				}

				bar.Increment()
				//fmt.Printf("Download: %v image: %v successful.\n", filePathName, url)

			}
			//color.Green("Run [%v], [%v]images successfully!!!\n", filePath, len(*availableImgs))
			picMapping[filePath] = picMap
			bar.Finish()
			wg.Done()

		}(filePath)
	}
	wg.Wait()
	color.Green("Successful handling of [%v] files.\n", len(*allFile))

	if err != nil {
		log.Fatal(err)
	}

	return &picMapping
}

// downLoadPath/filename-5cd1d11a5612f.jpg
func genFullFileName(downLoadPath, filePath string, url *string) *string {
	suffix := util.GetFileSuffix(*url)
	fullFileName := downLoadPath + "/" + util.GetFileSuffix(filePath) + "---" + suffix
	return &fullFileName
}
