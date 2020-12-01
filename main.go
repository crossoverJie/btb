package main

import (
	"btb/constants"
	"btb/io"
	"github.com/fatih/color"
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
			allFile, err := io.GetAllFile(markdownPath)
			if err != nil {
				log.Fatal("read file error")
			}

			for _, fileName := range allFile {
				color.Blue("fileName=[%v]", fileName)
			}
			color.Green("Run [%v] size successfully!!!", len(allFile))
			context, err := io.ReadFileContext(allFile[0])
			if err != nil {
				log.Fatal(err)
			}

			color.Red(context)

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
