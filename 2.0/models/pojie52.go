package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Pojie52Cookie struct {
	Cookies  string
	Url      string
	HtmlBody string
	PostForm url.Values
}

func (c *Pojie52Cookie) SetPojie52() {
	data := url.Values{}
	data.Set("mod", "task")
	data.Set("do", "apply")
	data.Set("id", "2")
	c.PostForm = data
	c.Url = "https://www.52pojie.cn/home.php?"
	err := c.GetPojie52()
	if err != nil {
		log.Println("吾爱破解论坛 签到失败")
		return
	}
	log.Println("吾爱破解论坛 已发送签到请求")
}

func (c *Pojie52Cookie) GetPojie52() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, ioutil.NopCloser(strings.NewReader(c.PostForm.Encode())))
	if err != nil {
		return err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //Post请求需要设置这个
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
	c.HtmlBody = string(body)
	return nil
}
