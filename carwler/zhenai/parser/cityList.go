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
			ParserFunc: ParseUserList, //ParseUserList,
		})
	}
	return result
}

const userListRe = `window.__INITIAL_STATE__=(.*?});`

//ParseUserList 用户List分析器
func ParseUserList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(userListRe)
	userMatches := re.FindAllSubmatch(contents, -1)
	profile, err := ParseProfile(userMatches[0][1])
	if err != nil {
		fmt.Printf("parseUserList userMatches fail :%s\n", err)
	}

	userList := profile.(AllProfile).MemberListData.MemberList
	result := engine.ParseResult{}
	for _, user := range userList {
		result.Items = append(result.Items, user)
	}
	result.Request = []engine.Request{}

	return result
}
