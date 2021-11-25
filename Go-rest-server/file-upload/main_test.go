package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUploadHandler(t *testing.T) {
	assert := assert.New(t)
	path := "/Users/yeonjemin/Desktop/lion.jpg"
	file, _ := os.Open(path)
	defer file.Close()

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("file_upload", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	defer writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/upload", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

}
