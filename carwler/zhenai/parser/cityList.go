package main

import (
	"regexp"
	"u2pppw/carwler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/.[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	mathes := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range mathes {
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParse,
		})
	}
	return result
}
