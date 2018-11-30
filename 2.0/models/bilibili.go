package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BiliCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
}

func (c *BiliCookie) SetBili() {
	c.Url = "https://api.live.bilibili.com/sign/doSign"
	err := c.GetBili()
	if err != nil {
		log.Println("签到Bilibili失败")
		return
	}
	var Bili struct {
		Msg string `json:"msg"`
	}
	json.Unmarshal(c.HtmlBody, &Bili) //[]Byte解析
	log.Println(" Bilibili 直播区	", Bili.Msg)

}

func (c *BiliCookie) GetBili() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, strings.NewReader("name=cjb"))
	if err != nil {
		return err
	}
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
