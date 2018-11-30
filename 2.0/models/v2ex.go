package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type V2exCookie struct {
	Cookies  string
	Url      string
	HtmlBody string
}

func (c *V2exCookie) SetV2ex() {
	c.Url = "https://www.v2ex.com/mission/daily"
	err := c.GetV2ex()
	if err != nil {
		log.Println("打开V2EX失败!")
		return
	}
	html := Regexp2(c.HtmlBody, `(<input type="button" class="super normal button" value="领取 X 铜币" onclick="location.href = '[^']+|每日登录奖励已领取)`)
	if html == `每日登录奖励已领取` {
		log.Println("V2ex	每日登录奖励已领取过")
	} else if html == "" {
		log.Println("V2EX	未登录")
	} else {
		c.Url = "https://www.v2ex.com" + Regexp2(html, `[^']+$`)
		err := c.GetV2ex()
		if err != nil {
			log.Println("V2EX	签到失败")
			return
		}
		log.Println("V2EX	签到成功")
	}

}

func (c *V2exCookie) GetV2ex() error {
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
	c.HtmlBody = string(body)
	return nil
}
