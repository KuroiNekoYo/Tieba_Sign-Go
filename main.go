package main

import (
	. "./TiebaSign"
	"container/list"
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
		likedTiebaList, err := GetLikedTiebaList()
		if err != nil {
			fmt.Println(err)
			return
		}
		linkedList := list.New() // Create sign list
		for _, tieba := range likedTiebaList {
			linkedList.PushBack(tieba)
		}
		for {
			listItem := linkedList.Front()
			if listItem == nil {
				break
			}
			linkedList.Remove(listItem)
			tieba := listItem.Value.(LikedTieba)
			status, message, exp := TiebaSign(tieba)
			fmt.Printf("%s\t%d: %s\tEXP+%d\n", ToUtf8(tieba.Name), status, message, exp)
			if exp > 0 || status == 1 {
				time.Sleep(1e9)
			}
			if status == 1 {
				linkedList.PushBack(tieba) // push failed items back to list
			}
		}
		time.Sleep(3e9)
	} else {
		time.Sleep(5e9)
	}
}
