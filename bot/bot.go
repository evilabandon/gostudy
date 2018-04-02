package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func get(url string) {
	fmt.Println(url)
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("err")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("body err")
	}
	fmt.Println(string(body))
}

func main()  {
	get("http://www.youku.com")

}