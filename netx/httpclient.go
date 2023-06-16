package netx

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ReqCallBack func(res *http.Response, body *[]byte)

func HttpGet(url string, cb ReqCallBack) error {
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("HttpGet Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

func HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) error {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		//fmt.Println("HttpPost Err:", err)
		return err
	}
	defer resp.Body.Close()
	return handleResponse(resp, cb)
}

func HttpPostString(url, body string, cb ReqCallBack) error {
	return HttpPost(url, "application/x-www-form-urlencoded", strings.NewReader(body), cb)
}

func HttpPostForm(url string, data url.Values, cb ReqCallBack) error {
	resp, err := http.PostForm(url, data)
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

//func Do(url, contentType, body string, cb ReqCallBack) {
//	client := &http.Client{}
//	req, errsx := http.NewRequest("POST", url, strings.NewReader(body))
//	if errsx != nil {
//		logx.Fatal(errsx)
//		return
//	}
//	req.Header.Set("Content-Type", contentType)
//	resp, errsx := client.Do(req)
//	defer resp.Body.Close()
//	handleResponse(resp, cb)
//}

func handleResponse(resp *http.Response, cb ReqCallBack) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("handleResponse Err:", err)
		return err
	}
	if nil == cb {
		return nil
	}
	cb(resp, &body)
	return nil
}
