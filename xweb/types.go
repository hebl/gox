package xweb

// 一些函数类型

import (
	"net/http"

	"golang.org/x/net/context"
)

//HTTPAPIFunc 函数类型
type HTTPAPIFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error

// FilterHanderFunc 权限过滤中间件
/*
 */
type FilterHanderFunc func(permission int, handler HTTPAPIFunc) HTTPAPIFunc

// Middleware 中间件函数
type Middleware func(handler HTTPAPIFunc) HTTPAPIFunc
