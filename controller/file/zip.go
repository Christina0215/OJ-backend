package file

import (
	"archive/zip"
	_ "archive/zip"
	_ "fmt"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	_ "github.com/satori/go.uuid"
	"io"
	"os"
	_ "strings"
	_ "time"
)

func GetZip(c *gin.Context)  {
	
}

func UnzipByPath(tarFile, mimeType string ,filename string) error {
	srcFile, err := os.Open(tarFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	return Unzip(srcFile, mimeType, filename)
}

func Unzip(srcFile *os.File, mimeType string, filename string) error {
	var enc mahonia.Decoder
	enc = mahonia.NewDecoder("gbk")
	zipFile, err := zip.OpenReader(srcFile.Name())
	if err != nil {
		return err
	}
	defer zipFile.Close()
	for _, innerFile := range zipFile.File {
		info := innerFile.FileInfo()
		if info.IsDir() {
			err = os.MkdirAll("public/problem/", os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}
		srcFile, err := innerFile.Open()
		if err != nil {
			return err
		}
		newFile, err := os.Create("public/problem/"+enc.ConvertString(innerFile.Name))
		if err != nil {
			return err
		}
		_, _ = io.Copy(newFile, srcFile)
		_ = newFile.Close()
	}
	defer srcFile.Close()
	return err
}
