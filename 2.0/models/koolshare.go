package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type KoolshareCookie struct {
	Cookies string
	Url     string
}

func (c *KoolshareCookie) SetKoolshare() {
	c.Url = "http://koolshare.cn/forum.php"
	err := c.GetKoolshare()
	if err != nil {
		log.Println("Koolshare 站点打开失败")
		return
	}
	log.Println("Koolshare 已发送签到请求")
}
func (c *KoolshareCookie) GetKoolshare() error {
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
	_ = body
	return nil
}
