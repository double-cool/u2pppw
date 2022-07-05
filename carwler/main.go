package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get(
		"http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("closer resp body fail :%s\n", err)
			return
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code", resp.StatusCode)
		return
	}
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Printf("%s\n", all)

}
