package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HdpdansCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
	Referers string
}

func (c *HdpdansCookie) SetHdpdans() {
	c.Url = "http://www.hdpfans.com/qiandao"
	c.Referers = "http://www.hdpfans.com/"
	err := c.GetHdpdans()
	if err != nil {
		log.Println("高清范论坛打开失败")
		return
	}
	formhash := Regexp3(string(c.HtmlBody), `type="hidden" name="formhash" value="([^"]+)`)
	c.Url = "http://www.hdpfans.com/plugin.php?id=k_misign:sign&operation=qiandao&formhash=" + formhash + "&format=empty&inajax=1&ajaxtarget=JD_sign"
	err = c.GetHdpdans()
	if err != nil {
		log.Println("高清范论坛签到失败")
		return
	}
	log.Println("高清范论坛	已发送签到请求")
}
func (c *HdpdansCookie) GetHdpdans() error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.Url, strings.NewReader("name=cjb"))
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("Referer", c.Referers)
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
