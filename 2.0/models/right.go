package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type RightCookie struct {
	Cookies string
	Url     string
}

func (c *RightCookie) SetRight() {
	c.Url = "https://www.right.com.cn/forum/forum.php"
	err := c.GetRight()
	if err != nil {
		log.Println("恩山论坛 站点打开失败")
		return
	}
	log.Println("恩山论坛 已发送签到请求")
}
func (c *RightCookie) GetRight() error {
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
