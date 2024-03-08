package utils

// 发起http请求
type HttpClient interface {
	//参数后面再说
	Post(url string)
	Get(url string)
}
