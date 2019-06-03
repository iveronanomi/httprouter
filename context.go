package httprouter

import (
	"context"
	"net/http"
	"reflect"
	"runtime"
)

type key int

const (
	// CtxHandleNameKey context handle name key
	CtxHandleNameKey key = iota
	// CtxPath context path key
	CtxPath
)

// string name of handle
func name(handle Handle) (name string) {
	name = runtime.FuncForPC(reflect.ValueOf(handle).Pointer()).Name()
	for i := len(name); i > 0; i-- {
		if name[i-1]^0x2e == 0 {
			return name[i:]
		}
	}
	return ""
}

// upgradeCtx upgrade context in request
func upgradeCtx(r *http.Request, n *node) *http.Request {
	*r = *r.WithContext(context.WithValue(context.WithValue(
		r.Context(), CtxPath, n.pathFull), CtxHandleNameKey, n.handleName))
	return r
}
