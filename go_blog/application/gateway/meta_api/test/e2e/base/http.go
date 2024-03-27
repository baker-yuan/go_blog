package base

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/stretchr/testify/assert"
)

type UploadFile struct {
	Name     string
	Filepath string
}

func HttpGet(url string, headers map[string]string) ([]byte, int, error) {
	return httpRequest(http.MethodGet, url, headers, "")
}

func HttpDelete(url string, headers map[string]string) ([]byte, int, error) {
	return httpRequest(http.MethodDelete, url, headers, "")
}

func HttpPut(url string, headers map[string]string, reqBody string) ([]byte, int, error) {
	return httpRequest(http.MethodPut, url, headers, reqBody)
}

func HttpPost(url string, headers map[string]string, reqBody string) ([]byte, int, error) {
	return httpRequest(http.MethodPost, url, headers, reqBody)
}

func httpRequest(method, url string, headers map[string]string, reqBody string) ([]byte, int, error) {
	var requestBody = new(bytes.Buffer)
	if reqBody != "" {
		requestBody = bytes.NewBuffer([]byte(reqBody))
	}
	req, err := http.NewRequest(method, url, requestBody)

	if err != nil {
		return nil, 0, err
	}
	req.Close = true

	// set header
	for key, val := range headers {
		req.Header.Add(key, val)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}

	return body, resp.StatusCode, nil
}

func BatchTestServerPort(times int, headers map[string]string, queryString string) map[string]int {
	t := getTestingHandle()
	url := APISIXSingleWorkerHost + "/server_port"
	if queryString != "" {
		url = url + "?" + queryString
	}

	res := map[string]int{}
	for i := 0; i < times; i++ {
		bodyByte, status, err := HttpGet(url, headers)
		assert.Nil(t, err)
		assert.Equal(t, 200, status)

		body := string(bodyByte)
		if _, ok := res[body]; !ok {
			res[body] = 1
		} else {
			res[body] += 1
		}
	}

	return res
}

func GetReader(reqParams map[string]string, contentType string, files []UploadFile) (io.Reader, string, error) {
	if strings.Index(contentType, "json") > -1 {
		bytesData, _ := json.Marshal(reqParams)
		return bytes.NewReader(bytesData), contentType, nil
	}
	if files != nil {
		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		for _, uploadFile := range files {
			file, err := os.Open(uploadFile.Filepath)
			if err != nil {
				return nil, "", err
			}
			defer file.Close()
			part, err := writer.CreateFormFile(uploadFile.Name, filepath.Base(uploadFile.Filepath))
			if err != nil {
				return nil, "", err
			}
			_, err = io.Copy(part, file)
		}
		for k, v := range reqParams {
			if err := writer.WriteField(k, v); err != nil {
				return nil, "", err
			}
		}
		if err := writer.Close(); err != nil {
			return nil, "", err
		}
		return body, writer.FormDataContentType(), nil
	}

	urlValues := url.Values{}
	for key, val := range reqParams {
		urlValues.Set(key, val)
	}

	reqBody := urlValues.Encode()

	return strings.NewReader(reqBody), contentType, nil
}
