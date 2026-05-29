package httpx

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// ReqCallBack is invoked with the HTTP response and response body bytes.
// ReqCallBack 在收到 HTTP 响应时被调用，并传入响应体字节。
type ReqCallBack func(res *http.Response, body *[]byte)

const (
	// DefaultTimeout default request timeout (20 seconds)
	// DefaultTimeout 默认20秒超时
	DefaultTimeout = 20 * time.Second // 默认20秒超时
)

var (
	DefaultClient = NewHttpClient()
)

// NewHttpClient creates an HttpClient with the default timeout.
// NewHttpClient 创建使用默认超时的 HttpClient。
func NewHttpClient() *HttpClient {
	return &HttpClient{Client: &http.Client{Timeout: DefaultTimeout}}
}

// HttpClient wraps http.Client with convenience request methods.
// HttpClient 封装 http.Client 并提供便捷的请求方法。
type HttpClient struct {
	Client *http.Client
}

// SetTimeout sets the request timeout.
// SetTimeout 设置超时时间。
func (o *HttpClient) SetTimeout(timeout time.Duration) {
	if nil == o.Client {
		return
	}
	o.Client.Timeout = timeout
}

// HttpGet sends a GET request and invokes cb with the response.
// HttpGet 发送 Get 请求。
func (o *HttpClient) HttpGet(url string, cb ReqCallBack) error {
	resp, err := o.Client.Get(url)
	if err != nil {
		//fmt.Println("HttpGet Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

// HttpPost sends a POST request with the given content type and body.
// HttpPost 发送 Post 请求。
func (o *HttpClient) HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) error {
	resp, err := o.Client.Post(url, contentType, body)
	if err != nil {
		//fmt.Println("HttpPost Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

// HttpPostForm sends a POST request with form-encoded data.
// HttpPostForm 发送 Post 表单请求。
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

// HttpPostString sends a POST request with a string body.
// HttpPostString 发送内容为字符串的 Post 请求。
func (o *HttpClient) HttpPostString(url, body string, cb ReqCallBack) error {
	return o.HttpPost(url, "application/x-www-form-urlencoded", strings.NewReader(body), cb)
}

// SetTimeout sets the timeout on the default HttpClient.
// SetTimeout 设置默认 HttpClient 的超时时间。
func SetTimeout(timeout time.Duration) {
	DefaultClient.SetTimeout(timeout)
}

// HttpGet sends a GET request using the default HttpClient.
// HttpGet 使用默认 HttpClient 发送 Get 请求。
func HttpGet(url string, cb ReqCallBack) error {
	return DefaultClient.HttpGet(url, cb)
}

// HttpPost sends a POST request using the default HttpClient.
// HttpPost 使用默认 HttpClient 发送 Post 请求。
func HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) error {
	return DefaultClient.HttpPost(url, contentType, body, cb)
}

// HttpPostString sends a POST request with a string body using the default HttpClient.
// HttpPostString 使用默认 HttpClient 发送内容为字符串的 Post 请求。
func HttpPostString(url, body string, cb ReqCallBack) error {
	return DefaultClient.HttpPostString(url, body, cb)
}

// HttpPostForm sends a POST form request using the default HttpClient.
// HttpPostForm 使用默认 HttpClient 发送 Post 表单请求。
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
