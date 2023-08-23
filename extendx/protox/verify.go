// Package protox
// Create on 2023/8/23
// @author xuzhuoxi
package protox

type FuncNewIReqVerify = func() IReqVerify

func NewIReqVerify() IReqVerify {
	return &ReqVerify{}
}

type IReqVerify interface {
	// Clear 清除
	Clear()
	// AppendVerifyHandler
	// 追加验证处理函数
	AppendVerifyHandler(handler FuncVerify)
	// AppendVerify
	// 追加验证处理器
	AppendVerify(verify IReqVerify)
	// Verify 验证请求入口
	Verify(name string, pid string, uid string) (rsCode int32)
}

type reqVerifyItem struct {
	handler FuncVerify
	verify  IReqVerify
}

type ReqVerify struct {
	handlers []*reqVerifyItem
}

func (o *ReqVerify) Clear() {
	o.handlers = nil
}

func (o *ReqVerify) AppendVerifyHandler(handler FuncVerify) {
	o.handlers = append(o.handlers, &reqVerifyItem{handler: handler})
}

func (o *ReqVerify) AppendVerify(verify IReqVerify) {
	o.handlers = append(o.handlers, &reqVerifyItem{verify: verify})
}

func (o *ReqVerify) Verify(name string, pid string, uid string) (rsCode int32) {
	if len(o.handlers) == 0 {
		return
	}
	for _, item := range o.handlers {
		if item == nil {
			continue
		}
		if item.handler != nil {
			code := item.handler(name, pid, uid)
			if code != CodeSuc {
				return code
			}
		} else if item.verify != nil {
			code := item.verify.Verify(name, pid, uid)
			if code != CodeSuc {
				return code
			}
		}
	}
	return
}
