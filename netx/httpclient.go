package netx

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type ReqCallBack func(res *http.Response, body *[]byte)

func HttpGet(url string, cb ReqCallBack) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	handleResponse(resp, cb)
}

func HttpPost(url, contentType string, body io.Reader, cb ReqCallBack) {
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	handleResponse(resp, cb)
}

func HttpPostString(url, body string, cb ReqCallBack) {
	HttpPost(url, "application/x-www-form-urlencoded", strings.NewReader(body), cb)
}

func HttpPostForm(url string, data url.Values, cb ReqCallBack) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()
	handleResponse(resp, cb)
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

func handleResponse(resp *http.Response, cb ReqCallBack) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(string(body))
	if nil != cb {
		cb(resp, &body)
	}
}
