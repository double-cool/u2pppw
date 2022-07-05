package main

import (
	"fmt"
	"regexp"
)

const text = `My email is 22213@qq.com
	emain is avc@asd.org
	 asd is aaa@aaa.log"
	 email3 is ddd@asd.com.cn`

func main() {
	re, err := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	if err != nil {
		return
	}
	mathch := re.FindAllStringSubmatch(text, -1)
	for _, m := range mathch {
		fmt.Println(m)
	}

}
