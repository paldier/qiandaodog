package models

import (
	"regexp"
	"strings"
)

func Regexp0(s1, s2 string) []string {
	r1 := regexp.MustCompile(s2)
	r2 := r1.FindAllString(s1, -1)
	return r2
}
func Regexp1(s1, s2, s3 string) string {
	r1 := regexp.MustCompile(s2)
	r2 := r1.ReplaceAllString(s1, s3)
	return r2
}
func Regexp2(s1, s2 string) string {
	r1 := regexp.MustCompile(s2)
	r2 := r1.FindString(s1)
	return r2
}
func Regexp3(s1, s2 string) string {
	r1 := regexp.MustCompile(s2)
	r2 := r1.FindSubmatch([]byte(s1))
	return string(r2[1])
}
func Replace0(s1, s2 string) string {
	return strings.Replace(s1, s2, "", -1)

}
