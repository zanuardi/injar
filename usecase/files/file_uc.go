package files

import (
	"context"
	"io"
	"mime/multipart"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type FileUC struct {
	fileRepository    Repository
	Context           context.Context
	DB                *gorm.DB
	webinarRepository Repository
	contextTimeout    time.Duration
}

func NewFileUC(timeout time.Duration, wr Repository) Usecase {
	return &FileUC{
		contextTimeout:    timeout,
		webinarRepository: wr,
	}
}

func (uc *FileUC) Store(fileType string, file *multipart.FileHeader) (Domain, error) {
	fileRepo := uc.fileRepository

	// source
	src, err := file.Open()
	if err != nil {
		log.Warn()
		return Domain{}, err
	}
	defer src.Close()

	folderPath := "public/" + fileType
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		log.Warn()
		return Domain{}, err
	}

	// destination
	dest := folderPath + "/" + file.Filename
	dst, err := os.Create(dest)
	if err != nil {
		log.Warn()
		return Domain{}, err
	}
	defer dst.Close()

	// copy file to folder
	if _, err = io.Copy(dst, src); err != nil {
		log.Warn()
		return Domain{}, err
	}

	fileEntity := Domain{
		Type: fileType,
		Path: dest,
	}

	fileEntity, err = fileRepo.Store(&fileEntity)
	if err != nil {
		log.Warn()
		return Domain{}, err
	}

	return Domain{}, err
}
