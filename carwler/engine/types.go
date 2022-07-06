package engine // Package engine 工具包

// Request 请求的结构体，包含请求url 和 对url中内容的处理
type Request struct {
	Url        string                   // 请求的url
	ParserFunc func([]byte) ParseResult // 分析函数
}

// ParseResult 分析结果的结构体
type ParseResult struct {
	Request []Request     // 当前分析结果中包含的请求结构体
	Items   []interface{} // 当前分析结果得到的内容
}

// NilParse 一个空的处理函数
func NilParse([]byte) ParseResult {
	return ParseResult{}
}
