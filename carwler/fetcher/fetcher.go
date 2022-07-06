package fetcher // Package fetcher 读取包

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// 读取网页的内容
func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("wrong staus code：%d", resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// <a href="http://www.zhenai.com/zhenghun/zhangjiajie" data-v-1573aa7c>张家界</a>
	return all, nil
}
