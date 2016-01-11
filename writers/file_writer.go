package writers

import (
	"os"
	"path/filepath"
)

type FileWriter struct {
	fng      FileNameGenerator
	file     *os.File
	fileName string
}

type FileNameGenerator func() string

func NewFileWriter(fileName FileNameGenerator) *FileWriter {
	return &FileWriter{
		fng: fileName,
	}
}

func (this *FileWriter) Write(data []byte) (err error) {
	err = this.checkFileName()
	if err != nil {
		return
	}
	_, err = this.file.Write(data)
	return err
}

func (this *FileWriter) Close() error {
	if this.file == nil {
		return nil
	}
	return this.file.Close()
}

func (this *FileWriter) checkFileName() error {
	newFileName := this.fng()
	if this.file == nil || this.fileName != newFileName {
		this.fileName = newFileName
		err := this.openFile()
		if err != nil {
			return err
		}
	}
	return nil
}

func (this *FileWriter) openFile() (err error) {
	fullPath := this.fileName
	fullPath, err = filepath.Abs(fullPath)
	if err != nil {
		return
	}
	err = this.makeDir(fullPath)
	if err != nil {
		return
	}
	file, err := os.OpenFile(fullPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return
	}
	this.file = file
	return nil
}

func (this *FileWriter) makeDir(fullPath string) error {
	dir := filepath.Dir(fullPath)
	return os.MkdirAll(dir, os.ModePerm)
}

var _ Writer = &FileWriter{}
