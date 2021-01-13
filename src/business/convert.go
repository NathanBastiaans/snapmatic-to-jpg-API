package business

import (
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// RequestModel holds the data for the request
type RequestModel struct {
	Image *multipart.FileHeader `form:"image" binding:"required"`
}

// Converts handles the convert request
func Convert(c *gin.Context) {
	var data RequestModel

	if err := c.ShouldBind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set new filename
	var filename = data.Image.Filename + ".jpg"

	// Open the image
	file, err := data.Image.Open()
	if err != nil || file == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Close the file when done
	defer file.Close()

	// Remove first 292 bytes of data
	var offset int64 = 292

	if _, err = file.Seek(offset, 0); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Check if the file needs to be saved locally
	if c.DefaultQuery("save", "false") != "false" {
		if err = WriteToFile(filename, fileData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// Return the converted data
	c.JSON(http.StatusOK, gin.H{
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
