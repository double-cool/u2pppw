package main

import (
	"u2pppw/carwler/engine"
	"u2pppw/carwler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "https://www.zhenai.com/zhenghun",
		ParserFunc: profile.ParseCityList,
	})

	return

}
