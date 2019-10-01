package engine

import (
	"os"
	"path"
)

func GetFiles(fileNames []string) ([]*os.File, error) {
	var files []*os.File
	for _,fileName := range fileNames {
		currentDir, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		fullFilePath := path.Join(currentDir, fileName)
		file, err := os.Open(fullFilePath)
		if err != nil {
			return nil, err
		}

		files = append(files, file)
	}
	return files, nil
}
