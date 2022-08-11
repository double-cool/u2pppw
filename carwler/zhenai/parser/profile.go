package profile

import (
	"encoding/json"
	"fmt"
)

type Member struct {
	NickName         string `json:"nickName"`
	Sex              int    `json:"sex"`
	Height           int    `json:"height"`
	Constellation    string `json:"constellation"`
	MemberID         int64  `json:"memberID"`
	Education        string `json:"education"`
	Marriage         string `json:"marriage"`
	Age              int    `json:"age"`
	Occupation       string `json:"occupation"`
	WorkCity         string `json:"workCity"`
	Salary           string `json:"salary"`
	IntroduceContent string `json:"introduceContent"`
}
type MemberListData struct {
	MemberList []Member `json:"memberList"`
}

type AllProfile struct {
	MemberListData MemberListData `json:"memberListData"`
}

func ParseProfile(profileJson []byte) (interface{}, error) {
	fmt.Printf("json  :%s\n", profileJson)
	var profileArr AllProfile
	err := json.Unmarshal(profileJson, &profileArr)
	if err != nil {
		fmt.Printf("json Unmarshal fail :%s\n", err)
		return nil, err
	}

	return profileArr, nil
}
