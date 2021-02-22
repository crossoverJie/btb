package main

import (
	"btb/constants"
	"btb/service"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var model, token string
	var uploadService service.UploadService
	downloadPath := constants.DownloadPath
	markdownPath := constants.MarkdownPath

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "model",
				Usage:       "Operating mode; r:replace, b:backup",
				DefaultText: "b",
				Aliases:     []string{"m"},
				Required:    true,
				Destination: &model,
			},
			&cli.StringFlag{
				Name:        "download-path",
				Usage:       "The path where the image is stored",
				Aliases:     []string{"dp"},
				Destination: &downloadPath,
				Required:    true,
				Value:       constants.DownloadPath,
			},
			&cli.StringFlag{
				Name:        "markdown-path",
				Usage:       "The path where the markdown file is stored",
				Aliases:     []string{"mp"},
				Destination: &markdownPath,
				Required:    true,
				Value:       constants.MarkdownPath,
			},
			&cli.StringFlag{
				Name:        "token",
				Usage:       "Upload token",
				Aliases:     []string{"tk"},
				Required:    true,
				Destination: &token,
				Value:       constants.MarkdownPath,
			},
		},
		Action: func(c *cli.Context) error {
			uploadService = &service.SMSUpload{
				Url:           "https://sm.ms/api/v2/upload",
				Authorization: token,
			}
			constants.AppModel = model
			color.Yellow("Current Model is [%v]", constants.Model[constants.AppModel])
			picMapping := service.DownLoadPic(markdownPath, downloadPath, uploadService)

			if model == constants.Replace {
				service.Replace(picMapping)
			}

			return nil
		},
		Name:  "btb",
		Usage: "Help you backup and replace your blog's images",
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
