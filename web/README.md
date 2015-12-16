## Web
   参考了`docker`的实现流程，`docker`提供了一个非常棒的范式。

### Example

```
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
```
