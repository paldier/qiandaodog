package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SmzdmCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
	Referers string
	PostForm url.Values
}

func (c *SmzdmCookie) SetSmzdm() {
	c.Url = "https://zhiyou.smzdm.com/user/checkin/jsonp_checkin"
	c.Referers = "https://www.smzdm.com/"
	err := c.GetSmzdm()
	if err != nil {
		log.Println("什么值得买	打开失败")
		return
	}
	var smzdm struct {
		Error_code int    `json:error_code`
		Error_msg  string `json:error_msg`
		Data       struct {
			Add_point int    `json:add_point`
			Slogan    string `json:slogan`
		} `json:"data"`
	}

	json.Unmarshal(c.HtmlBody, &smzdm) //[]Byte解析
	if smzdm.Error_code == 0 {
		log.Println("什么值得买	", Regexp1(smzdm.Data.Slogan, `<[^>]+>`, ""))
	} else {
		log.Println("什么值得买	", smzdm.Error_msg)
	}
}
func (c *SmzdmCookie) GetSmzdm() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, ioutil.NopCloser(strings.NewReader(c.PostForm.Encode())))
	if err != nil {
		return err
	}
	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //Post请求需要设置这个
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
