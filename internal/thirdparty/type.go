package thirdparty

import (
	"context"
	domain "github.com/Diana-Fox/xtimer/internal/domian"
)

// 后面考虑加一些失败策略等，所有接口放在这里
// 调用远程实际的
type Executor interface {
	//最后那个是附加信息，算预留的一些扩展参数内容，留给以后想扩展的实现使用
	Execute(ctx context.Context, method domain.MethodType,
		parameter string, url string, additional map[string]string) error
}
