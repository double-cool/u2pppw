package engine // Package engine 工具包

import (
	"log"
	"u2pppw/carwler/fetcher"
)

// Run 工具包的开始方法 在这里维护下请求队列
func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("fetching url is %s", r.Url)
		body, err := fetcher.Fetch(r.Url) // 取url
		if err != nil {
			log.Printf("Fetcher : error fetching url %s :%v", r.Url, err)
			continue
		}
		parseResult := r.ParserFunc(body) // 读取到链接的内容放入函数中处理

		requests = append(requests, parseResult.Request...) //把得到的所有url请求放入队列里 维护下这个爬虫队列

		// 输出得到的内容
		for _, item := range parseResult.Items {
			log.Printf("Got item is : %v", item) // %v表示不转义
		}

	}
}
