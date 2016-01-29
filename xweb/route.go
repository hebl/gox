package xweb

// 以下代码参考来自 docker

//Router router
type Router interface {
	Routes() []Route
}

// Route 路由表
type Route interface {
	Method() string
	Path() string
	Handler() HTTPAPIFunc
}

// localRoute defines an individual API route to connect with the docker daemon.
// It implements router.Route.
type localRoute struct {
	method  string
	path    string
	handler HTTPAPIFunc
}

// Method returns the http method that the route responds to.
func (l localRoute) Method() string {
	return l.method
}

// Path returns the subpath where the route responds to.
func (l localRoute) Path() string {
	return l.path
}

// Handler returns the APIFunc to let the server wrap it in middlewares
func (l localRoute) Handler() HTTPAPIFunc {
	return l.handler
}

// NewRoute initializes a new local router for the reouter
func NewRoute(method, path string, handler HTTPAPIFunc) Route {
	return localRoute{method, path, handler}
}

// NewGetRoute initializes a new route with the http method GET.
func NewGetRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("GET", path, handler)
}

// NewPostRoute initializes a new route with the http method POST.
func NewPostRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("POST", path, handler)
}

// NewPutRoute initializes a new route with the http method PUT.
func NewPutRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("PUT", path, handler)
}

// NewDeleteRoute initializes a new route with the http method DELETE.
func NewDeleteRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("DELETE", path, handler)
}

// NewOptionsRoute initializes a new route with the http method OPTIONS
func NewOptionsRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("OPTIONS", path, handler)
}

// NewHeadRoute initializes a new route with the http method HEAD.
func NewHeadRoute(path string, handler HTTPAPIFunc) Route {
	return NewRoute("HEAD", path, handler)
}
