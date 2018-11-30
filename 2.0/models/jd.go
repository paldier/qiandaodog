package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type JdCookie struct {
	Cookies  string
	Url      string
	HtmlBody []byte
	PostForm url.Values
}

func (c *JdCookie) SetJd() {
	data := make(url.Values)
	data.Set("reqData", "{}")
	data.Set("source", "jrm")
	c.PostForm = data
	c.Url = "https://ms.jr.jd.com/gw/generic/base/h5/m/baseSignInEncryptNew"
	err := c.GetJd()
	if err != nil {
		log.Println("打开京东金融失败")
		return
	}
	var animals struct {
		ResultCode int    `json:resultCode`
		ResultMsg  string `json:resultMsg`
		ResultData struct {
			ShowMsg string `json:showMsg`
		}
	}
	json.Unmarshal(c.HtmlBody, &animals)
	if animals.ResultCode != 0 {
		log.Println("京东签到:	", animals.ResultMsg)
		return
	}
	log.Println("京东钢镚:	", animals.ResultData.ShowMsg)
	c.Url = "https://api.m.jd.com/client.action?functionId=fBankSign&body=%7B%7D&client=ld&clientVersion=1.0.0&appId=ld"
	err = c.GetJdll()
	if err != nil {
		log.Println("打开京东个人中心失败")
		return
	}
	var animalsll struct {
		Code         string `json:code`
		ErrorMessage string `json:errorMessage`
	}
	json.Unmarshal(c.HtmlBody, &animalsll)
	if animalsll.Code != "0" {
		log.Println("京东流量:	领取失败")
		return
	}
	log.Println("京东流量:	", animalsll.ErrorMessage)

}
func (c *JdCookie) GetJd() error {
	client := &http.Client{}
	req, err := http.NewRequest("POST", c.Url, ioutil.NopCloser(strings.NewReader(c.PostForm.Encode())))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") //Post请求需要设置这个
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("Host", "ms.jr.jd.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
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
func (c *JdCookie) GetJdll() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, strings.NewReader("name=cjb"))
	if err != nil {
		return err
	}
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("Host", "api.m.jd.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")
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
