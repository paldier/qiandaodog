package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type BaiduCookie struct {
	Url      string
	Cookies  string
	Referers string
	HtmlBody string
}

func (c *BaiduCookie) SetBaidu() {
	c.Url = "https://wapp.baidu.com/mo/q----,sz@320_240-1-3---2/m?tn=bdFBW&tab=favorite"
	c.Referers = "https://wapp.baidu.com/"
	err := c.GetBaidu()
	if err != nil {
		log.Println("打开百度贴吧失败!")
		return
	}
	BaiduTiebaLink := Regexp0(c.HtmlBody, `>\d+\.<a href="[^"]+">[^<]+`)
	BaiduTiebaLinkLen := len(BaiduTiebaLink)
	if BaiduTiebaLinkLen == 0 {
		log.Println("未关注贴吧或百度Cookie错误")
		return
	}
	fmt.Println("--------------------------------正在签到百度贴吧：--------------------------------")
	LenTieba := 0
	CanTieba := 0
	for i := 0; i < BaiduTiebaLinkLen; i++ {
		BaiduTiebaLinkS := "https://" + Regexp3(BaiduTiebaLink[i], `"//([^"]+)`)
		BaiduTiebaName := Regexp2(BaiduTiebaLink[i], `[^>]+$`)
		c.Url = BaiduTiebaLinkS
		err := c.GetBaidu()
		if err != nil {
			log.Println("Get ", BaiduTiebaName, " 失败")
			continue
		}
		if strings.Contains(c.HtmlBody, ">签到<") == true {
			qiandao_url := "https://tieba.baidu.com" + strings.Replace(Regexp3(c.HtmlBody, `style="text-align:right;"><a href="([^"]+)`), `amp;`, "", -1) //筛选出签到链接
			c.Url = qiandao_url
			err := c.GetBaidu()
			if err != nil {
				log.Println(BaiduTiebaName, "	x")
				continue
			}
			log.Println(BaiduTiebaName, "	√")
			CanTieba++
			LenTieba++
		} else if strings.Contains(c.HtmlBody, ">已签到<") {
			LenTieba++
		} else {
			continue
		}
	}

	fmt.Println("--------------------------------百度贴吧签到完成,本次签到", CanTieba, "个吧,已签到", LenTieba, " / ", BaiduTiebaLinkLen)
	c.SetBaiduWenku()

}
func (c *BaiduCookie) SetBaiduWenku() {
	fmt.Println("--------------------------------正在签到百度文库：--------------------------------")
	c.Referers = "https://wenku.baidu.com/task/browse/daily"
	c.Url = "https://wenku.baidu.com/task/submit/signin"
	err := c.GetBaidu()
	if err != nil {
		log.Println("百度文库签到失败")
		return
	}
	c.Url = "https://tanbi.baidu.com/home/task/taskIndex"
	err = c.GetBaidu()
	if err != nil {
		log.Println("打开百度文库失败!")
		return
	}
	log.Println("百度文库已连续签到", Regexp3(c.HtmlBody, `<div class="countIcon">([^<]+)</div>`), "天")
}
func (c *BaiduCookie) GetBaidu() error {
	client := &http.Client{}
	req, err := http.NewRequest("GET", c.Url, strings.NewReader("name=cjb"))
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
	c.HtmlBody = string(body)
	return nil
}
