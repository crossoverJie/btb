package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type UploadService interface {
	upload(path string) (uploadPath string, err error)
}

type SMSUpload struct {
	Url           string
	Authorization string
}

func (s *SMSUpload) upload(path string) (uploadPath string, err error) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(path)
	defer file.Close()
	part1, err := writer.CreateFormFile("smfile", filepath.Base(path))
	_, err = io.Copy(part1, file)

	err = writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", s.Url, payload)

	if err != nil {
		return
	}
	req.Header.Add("Authorization", s.Authorization)
	req.Header.Add("Content-Type", "multipart/form-data")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	smsRes := struct {
		Success bool   `json:"success"`
		Code    string `json:"code"`
		Message string `json:"message"`
		Data    struct {
			URL string `json:"url"`
		} `json:"data"`
		RequestID string `json:"RequestId"`
	}{}
	err = json.Unmarshal(body, &smsRes)
	if err != nil {
		return
	}

	if smsRes.Code != "success" {
		log.Fatalf("Upload error msg=%v", smsRes.Message)
	}
	return smsRes.Data.URL, err
}

type MyUpload struct{}

func (s *MyUpload) upload(path string) (uploadPath string, err error) {
	return "", nil
}
