package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//engine.Run(engine.Request{
	//	Url:        "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})

	url := "https://xskydata.jobs.feishu.cn/school"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	// <a href="http://www.zhenai.com/zhenghun/zhangjiajie" data-v-1573aa7c>张家界</a>
	fmt.Printf("all:%s", all)
	// 把网页转为utf-8编码

	//
	//url := "https://xskydata.jobs.feishu.cn/school"
	//resp, err := http.Get(url)
	//if err != nil {
	//	return
	//}
	//
	//defer func(Body io.ReadCloser) {
	//	err := Body.Close()
	//	if err != nil {
	//		fmt.Printf("closer resp body fail :%s\n", err)
	//		return
	//	}
	//}(resp.Body)
	//
	//if resp.StatusCode != http.StatusOK {
	//	fmt.Println("Error: status code", resp.StatusCode)
	//	return
	//}
	//
	//all, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	return
	//}
	//// <a href="http://www.zhenai.com/zhenghun/zhangjiajie" data-v-1573aa7c>张家界</a>
	//fmt.Printf("all:%s", all)
	return

}
