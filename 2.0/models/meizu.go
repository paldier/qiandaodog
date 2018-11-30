package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type MeizuCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
}

func (c *MeizuCookie) SetMeizu() {
	c.Url = "https://bbs-act.meizu.cn/index.php?mod=signin&action=sign"
	err := c.GetMeizu()
	if err != nil {
		log.Println("魅族论坛打开失败!")
		return
	}
	var animals struct {
		Code    int    `json:code`
		Message string `json:message`
	}
	json.Unmarshal(c.HtmlBody, &animals)
	log.Println("魅族论坛", animals.Message)

}
func (c *MeizuCookie) GetMeizu() error {
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
