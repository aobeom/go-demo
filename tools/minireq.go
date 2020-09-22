package tools

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// MiniRequestBasic 提供基本 HTTP 请求
type MiniRequestBasic struct{}

// Minireq 初始化
var Minireq *MiniRequestBasic

func init() {
	Minireq = NewMiniReq()
}

// NewMiniReq 初始化 NewMiniReq
func NewMiniReq() (n *MiniRequestBasic) {
	n = new(MiniRequestBasic)
	return
}

// GetRes 获取 Response
//	Example:
//	headers := make(http.Header)
//	headers.Add("Content-Type", "application/json")
//	params := make(map[string]string)
//	params["key"] = value
func (mr *MiniRequestBasic) GetRes(url string, headers http.Header, params map[string]string) (res *http.Response) {
	httpClient := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(" [GetRes - Request Error]: ", err)
	}
	req.Header = headers
	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err = httpClient.Do(req)

	if err != nil {
		log.Panic(" [GetRes - Response Error]: ", err)
	}
	defer res.Body.Close()
	return
}

// GetBody Http Get 获取 Body 内容
//	Example:
//	headers := make(http.Header)
//	headers.Add("Content-Type", "application/json")
//	params := make(map[string]string)
//	params["key"] = value
func (mr *MiniRequestBasic) GetBody(url string, headers http.Header, params map[string]string) []byte {
	httpClient := http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(" [GetBody - Request Error]: ", err)
	}

	req.Header = headers

	if params != nil {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	res, err := httpClient.Do(req)

	if err != nil {
		log.Panic(" [GetBody - Response Error]: ", err)
	}

	if res.StatusCode != 200 {
		log.Panic(" [GetBody - Response Code != 200]: ", err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(" [GetBody - Body Error]: ", err)
	}
	return body
}

// PostBody Http Post 获取 Body 内容
//	Example:
//	headers := make(http.Header)
//	headers.Add("Content-Type", "application/json")
//	data := make(map[string]string)
//	data["key"] = value
func (mr *MiniRequestBasic) PostBody(url string, headers http.Header, data map[string]string) []byte {
	httpClient := http.Client{Timeout: 30 * time.Second}

	rawData, err := json.Marshal(data)
	if err != nil {
		log.Panic(" [Post - Data Error]: ", err)
	}
	reader := bytes.NewBuffer(rawData)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		log.Panic(" [Post - Request Error]: ", err)
	}

	username := ""
	password := ""

	req.Header = headers

	if headers.Get("username") != "" {
		username = headers.Get("username")
		req.Header.Del("username")
	}

	if headers.Get("password") != "" {
		password = headers.Get("password")
		req.Header.Del("password")
	}

	if username != "" && password != "" {
		req.SetBasicAuth(username, password)
	}

	res, err := httpClient.Do(req)

	if err != nil {
		log.Panic(" [Post - Response Error]: ", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(" [Get - Body Error]: ", err)
	}
	return body
}
