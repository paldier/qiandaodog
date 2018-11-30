package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type AcfunCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
}

func (c *AcfunCookie) SetAcfun() {
	c.Url = "http://www.acfun.cn/webapi/record/actions/signin"
	err := c.GetAcfun()
	if err != nil {
		log.Println("Acfunc	签到失败!")
		return
	}
	var animals struct {
		Message string `json:"message"`
	}
	json.Unmarshal(c.HtmlBody, &animals) //[]Byte解析
	log.Println("Acfun	", animals.Message)
}
func (c *AcfunCookie) GetAcfun() error {
	data := url.Values{}
	p := fmt.Sprintf("%d", (time.Now().UnixNano())/1e6)
	data.Set("channel", "0")
	data.Set("date", p)
	client := &http.Client{}

	req, err := http.NewRequest("POST", c.Url, ioutil.NopCloser(strings.NewReader(data.Encode())))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //Post请求需要设置这个
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.84 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	c.HtmlBody = body
	return nil
}
