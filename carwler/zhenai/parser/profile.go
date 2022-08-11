package profile

import (
	"encoding/json"
	"fmt"
)

type ObjectInfo struct {
	BasicInfo []string `json:"basicInfo"`
}

type AllProfile struct {
	ObjectInfo ObjectInfo `json:"objectInfo"`
}

func ParseProfile(profileJson []byte) (interface{}, error) {
	var profileArr AllProfile
	err := json.Unmarshal(profileJson, &profileArr)
	if err != nil {
		fmt.Printf("json Unmarshal fail :%s\n", err)
		return nil, err
	}

	return profileArr, nil
}

func ParseUserList(profileJson []byte) (AllProfile, error) {
	var profileArr AllProfile
	err := json.Unmarshal(profileJson, &profileArr)
	if err != nil {
		fmt.Printf("json Unmarshal fail :%s\n", err)
		return AllProfile{}, err
	}

	return profileArr, nil
}
