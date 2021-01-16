package business

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// RequestModel holds the data for the request
type RequestModel struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// Converts handles the convert request
func Convert(ctx *gin.Context) {
	var data RequestModel

	if err := ctx.ShouldBind(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set new filename
	var filename = data.Image.Filename + ".jpg"

	// Open the image
	file, err := data.Image.Open()
	if err != nil || file == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the file when done
	defer file.Close()

	// Remove first 292 bytes of data
	var offset int64 = 292

	if _, err = file.Seek(offset, 0); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the file needs to be saved locally
	if ctx.DefaultQuery("save", "false") != "false" {
		if err = WriteToFile(filename, fileData); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Return the converted data
	ctx.JSON(http.StatusOK, gin.H{
		"filename": filename,
		"filedata": fileData,
	})
}

// WriteToFile writes the given data to the filename
func WriteToFile(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer file.Close()

	if _, err = io.WriteString(file, string(data)); err != nil {
		return err
	}

	return file.Sync()
}
