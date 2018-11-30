package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type HostlocCookie struct {
	Cookies  string
	Url      string
	HtmlBody string
}

func (c *HostlocCookie) SetHostloc() {
	c.Url = "https://www.hostloc.com/forum.php"
	err := c.GetHostloc()
	if err != nil {
		log.Println("Hostloc	首页打开失败")
	}
	HostlocLen := []string{"25650", "7436", "22176", "23376", "132", "26477", "25285", "26532", "25728", "26440", "18756", "12368", "26564"}
	for i := 0; i < 12; i++ {
		c.Url = "https://www.hostloc.com/space-uid-" + HostlocLen[i] + ".html"
		err := c.GetHostloc()
		if err != nil {
			continue
		}
	}
	log.Println("Hostloc	已发送获取积分请求!")

}
func (c *HostlocCookie) GetHostloc() error {
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
