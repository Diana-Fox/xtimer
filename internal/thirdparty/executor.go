package thirdparty

import (
	"context"
	domain "github.com/Diana-Fox/xtimer/internal/domian"
	"github.com/Diana-Fox/xtimer/internal/utils"
)

// http执行器
type HttpExecutor struct {
	hc utils.HttpClient //因为是http请求
	l  utils.Logger     //打日志
}

func (h *HttpExecutor) Execute(ctx context.Context,
	method domain.MethodType, parameter string,
	url string, additional map[string]string) error {
	//todo 去用http发送请求,请求发送的具体组织暂时先放这，等后面再实现
	if method == domain.POST {
		h.hc.Post(url)
	} else {
		h.hc.Get(url)
	}

	//TODO implement me
	panic("implement me")
}
