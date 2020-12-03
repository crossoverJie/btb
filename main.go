package main

import (
	"btb/constants"
	"btb/service"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	var model string
	downloadPath := constants.DownloadPath
	markdownPath := constants.MarkdownPath

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "model",
				Usage:       "operating mode; r:replace, b:backup",
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
				Value:       constants.DownloadPath,
			},
			&cli.StringFlag{
				Name:        "markdown-path",
				Usage:       "The path where the markdown file is stored",
				Aliases:     []string{"mp"},
				Destination: &markdownPath,
				Value:       constants.MarkdownPath,
			},
		},
		Action: func(c *cli.Context) error {
			service.DownLoadPic(markdownPath)

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
