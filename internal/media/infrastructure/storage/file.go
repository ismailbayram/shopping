package infrastructure

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type FileStorage struct {
	baseDir string
}

func NewFileStorage(mediaRoot string) FileStorage {
	f, _ := os.Getwd()
	baseDir := filepath.Join(filepath.Dir(f), filepath.Base(f), mediaRoot)

	_, err := os.Stat(baseDir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(baseDir, 0755)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return FileStorage{
		baseDir: baseDir,
	}
}

func (ifs FileStorage) Upload(name string, content []byte) (string, error) {
	fileName := generateFileName(name)
	path := filepath.Join(ifs.baseDir, fileName)

	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		return "", err
	}
	if _, err = file.Write(content); err != nil {
		return "", err
	}

	return fileName, nil
}

func generateFileName(name string) string {
	splitted := strings.Split(name, ".")
	return fmt.Sprintf("%s_%s.%s", slug.Make(splitted[0]), uuid.New(), splitted[1])
}
