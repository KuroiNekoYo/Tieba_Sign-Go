package main

import (
	. "./TiebaSign"
	"fmt"
	"io/ioutil"
	"time"
)

func main() {
	var username, password string
	fmt.Print("Enter your Baidu ID: ")
	fmt.Scan(&username)
	if username == "" {
		return
	}
	fmt.Print("Enter your Baidu Password: ")
	fmt.Scan(&password)
	if password == "" {
		return
	}
	result, loginErr := BaiduLogin(username, password)
	if loginErr == nil && result > 0 {
		fmt.Println("Successfully login")
		cookieStr := ""
		for _, cookie := range GetCookies() {
			cookieStr += cookie.Name + "=" + cookie.Value + "\n"
		}
		ioutil.WriteFile("cookie.txt", []byte(cookieStr), 0644)

		fmt.Println("Your cookie has been written into cookie.txt")
		GetLikedTiebaList()
		time.Sleep(2e9)
	} else {
		time.Sleep(5e9)
	}
}