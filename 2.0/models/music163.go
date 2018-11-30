package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Music163Cookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
}

func (c *Music163Cookie) SetMusic163() {
	c.Url = "http://music.163.com/api/point/dailyTask"
	err := c.GetMusic163()
	if err != nil {
		log.Println("网易云音乐打开失败")
		return
	}
	var music163 struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	json.Unmarshal(c.HtmlBody, &music163) //[]Byte解析
	if music163.Code == 301 {
		log.Println("网易云音乐登录失败")
	} else {
		log.Println("网易云音乐签到", music163.Msg)
	}
}
func (c *Music163Cookie) GetMusic163() error {
	data := url.Values{}
	data.Set("type", "1")
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
