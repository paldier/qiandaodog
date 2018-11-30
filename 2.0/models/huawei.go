package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type HuaweiCookie struct {
	Cookies  string
	Url      string
	HtmlBody string
	PostForm url.Values
}

func (c *HuaweiCookie) SetHuawei() {
	c.Url = "https://club.huawei.com/dsu_paulsign-sign.html"
	err := c.GetHuawei()
	if err != nil {
		log.Println("花粉俱乐部打开失败!")
		return
	}

	formhash := Regexp3(c.HtmlBody, `name="formhash" value="([^"]+)`)
	data := url.Values{}
	data.Set("operation", "qiandao")
	data.Set("formhash", formhash)
	c.PostForm = data
	c.Url = "https://club.huawei.com/plugin.php?id=dsu_paulsign:sign"
	err = c.PostHuawei()
	if err != nil {
		log.Println("花粉俱乐部签到失败")
		return
	}
	var animals struct {
		Credit int    `json:"credit"`
		Url    string `json:url`
	}
	fmt.Println(c.HtmlBody)
	json.Unmarshal([]byte(c.HtmlBody), &animals) //[]Byte解析
	if animals.Credit != 0 {
		log.Println("花粉俱乐部签到成功")
	} else {
		log.Println("花粉俱乐部	签到失败")
	}

}
func (c *HuaweiCookie) GetHuawei() error {
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
func (c *HuaweiCookie) PostHuawei() error {
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
