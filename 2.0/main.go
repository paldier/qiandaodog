/*
当前版本：2.1
*/

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"GO/qiandaodog/2.1/models"
)

func Openfile() (cookie []string) {
	fi, err := os.Open("cookie.txt") //读取文件夹里的cookie.txt
	if err != nil {
		log.Println("No cookie.txt file")
	}
	defer fi.Close()
	rd := bufio.NewReader(fi)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		cookie = append(cookie, strings.Replace(line, "\n", "", -1))
		if err != nil || io.EOF == err {
			break
		}
	}
	return
}
func StringSplitn(s1, s2 string) []string {
	r1 := strings.SplitAfterN(s1, s2, 10)
	r1[1] = strings.Replace(r1[1], "\r", "", -1)
	r1[1] = strings.Replace(r1[1], "\n", "", -1)
	return r1
}
func main() {
	txtbody := Openfile()
	t := len(txtbody)
	for i := 0; i < t; i++ {
		CookieBool := StringSplitn(txtbody[i], `"=`)
		if CookieBool[1] == "" {
			continue
		}
		switch {
		case CookieBool[0] == `"baidu"=`: //百度贴吧及文库
			c := &models.BaiduCookie{}
			c.Cookies = CookieBool[1]
			c.SetBaidu()
			break
		case CookieBool[0] == `"v2ex"=`: //V2EX
			c := &models.V2exCookie{}
			c.Cookies = CookieBool[1]
			c.SetV2ex()
			break
		case CookieBool[0] == `"hostloc"=`: //hostloc
			c := &models.HostlocCookie{}
			c.Cookies = CookieBool[1]
			c.SetHostloc()
			break

		case CookieBool[0] == `"acfun"=`: //A站
			c := &models.AcfunCookie{}
			c.Cookies = CookieBool[1]
			c.SetAcfun()
			break

		case CookieBool[0] == `"bilibili"=`: //哔哩哔哩直播区
			c := &models.BiliCookie{}
			c.Cookies = CookieBool[1]
			c.SetBili()
			break
		case CookieBool[0] == `"163music"=`: //网易云音乐
			c := &models.Music163Cookie{}
			c.Cookies = CookieBool[1]
			c.SetMusic163()
			break

		case CookieBool[0] == `"miui"=`: //miui论坛
			c := &models.MiuiCookie{}
			c.Cookies = CookieBool[1]
			c.SetMiui()
			break

		case CookieBool[0] == `"52pojie"=`: //吾爱破解
			c := &models.Pojie52Cookie{}
			c.Cookies = CookieBool[1]
			c.SetPojie52()
			break

		case CookieBool[0] == `"kafan"=`: //卡饭
			c := &models.KafanCookie{}
			c.Cookies = CookieBool[1]
			c.SetKafan()
			break

		case CookieBool[0] == `"smzdm"=`: //什么值得买
			c := &models.SmzdmCookie{}
			c.Cookies = CookieBool[1]
			c.SetSmzdm()
			break

		case CookieBool[0] == `"jd"=`: //京东领钢镚
			c := &models.JdCookie{}
			c.Cookies = CookieBool[1]
			c.SetJd()
			break

		case CookieBool[0] == `"zimuzu"=`: //人人字幕组
			c := &models.ZimuzuCookie{}
			c.Cookies = CookieBool[1]
			c.SetZimuzu()
			break
		case CookieBool[0] == `"gztown"=`: //港知堂社区PT站
			c := &models.GztownCookie{}
			c.Cookies = CookieBool[1]
			c.SetGztown()
			break
		case CookieBool[0] == `"meizu"=`: //魅族论坛
			c := &models.MeizuCookie{}
			c.Cookies = CookieBool[1]
			c.SetMeizu()
			break

		case CookieBool[0] == `"hdpfans"=`: //高清范论坛
			c := &models.HdpdansCookie{}
			c.Cookies = CookieBool[1]
			c.SetHdpdans()
			break

		case CookieBool[0] == `"chh"=`: //CHH
			c := &models.ChhCookie{}
			c.Cookies = CookieBool[1]
			c.SetChh()
			break

		case CookieBool[0] == `"koolshare"=`: //Koolshare
			c := &models.KoolshareCookie{}
			c.Cookies = CookieBool[1]
			c.SetKoolshare()
			break

		case CookieBool[0] == `"right"=`: //恩山
			c := &models.RightCookie{}
			c.Cookies = CookieBool[1]
			c.SetRight()
			break

		case CookieBool[0] == `"huawei"=`: //花粉俱乐部
			c := &models.HuaweiCookie{}
			c.Cookies = CookieBool[1]
			c.SetHuawei()
			break

		default:
			break

		}

	}
}
