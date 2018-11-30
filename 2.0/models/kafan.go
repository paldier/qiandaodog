package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type KafanCookie struct {
	Cookies  string
	Url      string
	HtmlBody string
	PostForm url.Values
}

func (c *KafanCookie) SetKafan() {
	c.Url = "https://bbs.kafan.cn/"
	err := c.GetKafan()
	if err != nil {
		log.Println("卡饭论坛打开失败!")
		return
	}
	formhash := Regexp3(c.HtmlBody, `name="formhash" value="([^"]+)`)
	data := url.Values{}
	data.Set("id", "dsu_amupper")
	data.Set("ppersubmit", "true")
	data.Set("formhash", formhash)
	data.Set("infloat", "yes")
	data.Set("handlekey", "dsu_amupper")
	data.Set("inajax", "1")
	data.Set("ajaxtarget", "fwin_content_dsu_amupper")
	c.PostForm = data
	c.Url = "https://bbs.kafan.cn/plugin.php"
	err = c.PostKafan()
	if err != nil {
		log.Println("卡饭论坛签到失败!")
		return
	}
	log.Println("卡饭论坛	已发送签到请求")
}
func (c *KafanCookie) GetKafan() error {
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
func (c *KafanCookie) PostKafan() error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.Url, ioutil.NopCloser(strings.NewReader(c.PostForm.Encode())))
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
	c.HtmlBody = string(body)
	return nil
}
