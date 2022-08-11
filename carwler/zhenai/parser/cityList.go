package profile

import (
	"fmt"
	"regexp"
	"u2pppw/carwler/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/.[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市分析器
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matchs {
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParse, //ParseUserList,
		})
	}
	return result
}

const userListRe = `window.__INITIAL_STATE__=(.*?);`

//ParseUserList 用户List分析器
func ParseUserList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userListRe)
	userMatches := re.FindAllSubmatch(contents, -1)
	fmt.Printf("aaa %s", userMatches)

	userList, err := parseUserList(userMatches[0][1])
	if err != nil {
		fmt.Printf("parseUserList userMatches fail :%s\n", err)
	}
	result := engine.ParseResult{}
	for _, m := range userList {
		result.Items = append(result.Items, string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParse,
		})
	}
	return engine.ParseResult{}
}

// ParseUser 用户分析器
const userRe = `window.__INITIAL_STATE__=(.*?});`

func ParseUser(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userRe)
	userMatches := re.FindAllSubmatch(contents, -1)

	user, err := ParseProfile(userMatches[0][1])
	if err != nil {
		fmt.Printf("parseUserList userMatches fail :%s\n", err)
	}

	result := engine.ParseResult{}
	result.Items = append(result.Items, user)
	result.Request = []engine.Request{}
	return result

}
