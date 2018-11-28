package gdc

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

// PostFile 文件上传
func PostFile(filename, targetURL string) (statusCode int, body []byte, err error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	// 创建FormFile
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		return
	}
	// 打开文件
	fh, err := os.Open(filename)
	if err != nil {
		return
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return
	}
	contentType := bodyWriter.FormDataContentType()
	if s, err := fh.Stat(); err == nil {
		bodyWriter.WriteField("size", fmt.Sprintf("%d", s.Size()))
	}
	bodyWriter.Close()
	resp, err := http.Post(targetURL, contentType, bodyBuf)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	statusCode = resp.StatusCode
	body, err = ioutil.ReadAll(resp.Body)
	return
}
