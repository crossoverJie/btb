package service

import (
	"btb/util"
	"log"
	"os"
	"strings"
)

func Replace(picMapping *map[string]map[string]string) {
	for filePath, picMap := range *picMapping {

		// Temp map used for mark replace text.
		var replaceMap = make(map[int]string)

		allLine, err := util.ReadFileLine(filePath)
		if err != nil {
			continue
		}
		file, err := os.Create(filePath)
		if err != nil {
			log.Fatalf("Rewrite [%s] file error=[%v]", filePath, err)
		}
		defer file.Close()

		// Mark place
		for index, line := range *allLine {
			for oldPic, _ := range picMap {

				if strings.Contains(line, oldPic) {
					replaceMap[index] = oldPic
				}
			}

		}

		// Loop for replacing text.
		for index, line := range *allLine {
			if len(replaceMap[index]) > 0 {
				line = strings.ReplaceAll(line, replaceMap[index], picMap[replaceMap[index]])
			}
			_, err := file.Write([]byte(line + "\n"))
			if err != nil {
				log.Fatalf("Rewrite [%s] file error=[%v]", filePath, err)
			}
		}

	}
}
