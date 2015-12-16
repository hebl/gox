package web

// 以下代码参考来自 docker

//Router router
type Router interface {
	Routes() []Route
}

// Route 路由表
type Route struct {
	Path    string
	Method  string
	Handler HTTPAPIFunc
}

// NewRoute initializes a new local router for the reouter
func NewRoute(method, path string, handler HTTPAPIFunc) Route {
	return Route{method, path, handler}
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

/**
一个模式

type MyRouter struct{
    routes []web.Route
}

func (r *myrouter) Routes() []web.Route {
	return r.routes
}

func NewMyRouter() web.Router{
    r := &MyRouter{}

    r.initRoutes()

    return r
}

func (r *myrouter) initRoutes(){
    r.routes = []web.Route{
        //GET
        web.NewGetRoute("/get", r.getHandler),
        //...
    }
}

// server

main:
   router := NewMyRouter()
   server := web.New(configFile)
   server.Init(router)
   server.StartServer()
*/
