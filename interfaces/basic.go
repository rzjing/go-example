/*
@ Author:       Wang Xiaoqiang <https://github.com/rzjing>
@ File:         basic.go
@ Create Time:  2020/5/3 10:43
@ Software:     GoLand
*/

package interfaces

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type Params struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	Offset   int    `json:"-" form:"-"`
	Sort     string `json:"sort" form:"sort"`
}

func (p *Params) Init() {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.PageSize <= 0 || p.PageSize > 100 {
		p.PageSize = 10
	}

	if p.Sort == "" {
		p.Sort = "updated_at desc"
	}

	p.Offset = (p.Page - 1) * p.PageSize
}

func Client(method, url string) (buf []byte, err error) {
	client := &http.Client{}
	defer client.CloseIdleConnections()

	var req *http.Request
	switch strings.ToUpper(method) {
	case "GET":
		req, err = http.NewRequest("GET", url, nil)
	default:
		return
	}

	if err != nil {
		return
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.129 Safari/537.36")

	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		return nil, errors.New("http response status code is: " + resp.Status)
	}
	defer resp.Body.Close()

	buf, err = ioutil.ReadAll(resp.Body)
	return
}
