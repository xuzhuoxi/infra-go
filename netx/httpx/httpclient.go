package httpx

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ReqCallBack func(res *http.Response, body *[]byte)

const (
	DefaultTimeout = 20 * time.Second // 默认20秒超时
)

var (
	DefaultClient = NewHttpClient()
)

func NewHttpClient() *HttpClient {
	return &HttpClient{Client: &http.Client{Timeout: DefaultTimeout}}
}

type HttpClient struct {
	Client *http.Client
}

// SetTimeout
// 设置超时时间
func (o *HttpClient) SetTimeout(timeout time.Duration) {
	if nil == o.Client {
		return
	}
	o.Client.Timeout = timeout
}

// HttpGet
// 发送Get请求
func (o *HttpClient) HttpGet(url string, cb ReqCallBack) error {
	resp, err := o.Client.Get(url)
	if err != nil {
		//fmt.Println("HttpGet Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

// HttpPost
// 发送Post请求
func (o *HttpClient) HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) error {
	resp, err := o.Client.Post(url, contentType, body)
	if err != nil {
		//fmt.Println("HttpPost Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

// HttpPostForm
// 发送Post请求
func (o *HttpClient) HttpPostForm(url string, data url.Values, cb ReqCallBack) error {
	resp, err := o.Client.PostForm(url, data)
	if err != nil {
		//fmt.Println("HttpPostForm Err:", err)
		return err
	}
	defer resp.Body.Close()
	err = handleResponse(resp, cb)
	if err != nil {
		return err
	}
	return nil
}

// HttpPostString
// 发送Post请求
// 内容为字符串
func (o *HttpClient) HttpPostString(url, body string, cb ReqCallBack) error {
	return o.HttpPost(url, "application/x-www-form-urlencoded", strings.NewReader(body), cb)
}

// SetTimeout
// 设置超时时间
// 使用默认的HttpClient
func SetTimeout(timeout time.Duration) {
	DefaultClient.SetTimeout(timeout)
}

// HttpGet
// 发送Get请求
// 使用默认的HttpClient
func HttpGet(url string, cb ReqCallBack) error {
	return DefaultClient.HttpGet(url, cb)
}

// HttpPost
// 发送Post请求
// 使用默认的HttpClient
func HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) error {
	return DefaultClient.HttpPost(url, contentType, body, cb)
}

// HttpPostString
// 发送Post请求
// 内容为字符串
// 使用默认的HttpClient
func HttpPostString(url, body string, cb ReqCallBack) error {
	return DefaultClient.HttpPostString(url, body, cb)
}

// HttpPostForm
// 发送Post请求
// 使用默认的HttpClient
func HttpPostForm(url string, data url.Values, cb ReqCallBack) error {
	return DefaultClient.HttpPostForm(url, data, cb)
}

func handleResponse(resp *http.Response, cb ReqCallBack) error {
	if nil == cb {
		return nil
	}
	code := resp.StatusCode
	// 2xx 状态码处理
	if code >= http.StatusOK && code < http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		cb(resp, &body)
		return nil
	} else {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		cb(resp, &body)
		return nil
	}
}
