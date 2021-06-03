package utils

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Sha256Check(target, origin string) bool {
	check := sha256.New()
	check.Write([]byte(origin))
	return hex.EncodeToString(check.Sum([]byte(""))) == target
}

func Sha256Get(origin string) string {
	target := sha256.New()
	target.Write([]byte(origin))
	return hex.EncodeToString(target.Sum([]byte("")))
}

func Unzip(src, dest, filename string) (err error, num int) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err, 0
	}
	defer func() {
		r.Close()
		err = os.Remove(src)
	}()
	for _, f := range r.File {
		fpath := filepath.Join(dest, filename, strings.Split(f.Name, "/")[0])
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("%s: illegal file path", fpath), 0
		}
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(filepath.Join(dest, filename), os.ModePerm)
			continue
		}
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err, 0
		}
		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err, 0
		}
		rc, err := f.Open()
		if err != nil {
			return err, 0
		}
		_, err = io.Copy(outFile, rc)
		_ = outFile.Close()
		_ = rc.Close()
		if err != nil {
			return err, 0
		}
		num++
	}
	return err, num / 2
}
