package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"io"
	"os"
	"strings"
)

func Upload(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
	}
	mimeType := strings.Split(header.Filename, ".")[1]
	filename := uuid.NewV4().String()
	fmt.Println(mimeType)
	var out *os.File
	switch mimeType {
	case "jpg", "png", "jpeg":
		out, err = os.Create("public/image/" + filename + "." + mimeType)
	case "zip":
		out, err = os.Create("public/temp/" + filename + "." + mimeType)
	default:
		return
	}
	if err != nil {
		c.JSON(403, gin.H{"message": err.Error()})
	}
	fmt.Println(out)
	if _, err = io.Copy(out, file); err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"filename": filename + "." + mimeType})
	defer func() {
		if err := out.Close(); err != nil {
			panic(err)
		}
	}()
}
