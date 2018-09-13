package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println("123.12.11.111", IsIp("123.12.11.111"))
}

func IsIp(ip string) bool {
	if m, _ := regexp.MatchString("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}", ip); m {
		return m
	}

	return true
}

func Matchd(str string) bool {
	matched, _ := regexp.Match("^[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}.[0-9]{1,3}", []byte(str))
	return matched
}

func trans(src string) string {
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	return re.ReplaceAllStringFunc(src, strings.ToLower)
}
