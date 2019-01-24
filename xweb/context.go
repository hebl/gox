package xweb

import "net/http"

//Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

// Params is a Param-slice, as returned by the router. The slice is ordered, the first URL parameter is also the first slice value. It is therefore safe to read values by the index.
type Params []Param

//Context 上下文信息
type Context struct {
	Request *http.Request
	Writer  http.ResponseWriter

	Params   Params
	handlers HandlersChain
	Errors   errorMsgs
	Accepted []string
}
