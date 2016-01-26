package main

import (
	"net/http"

	"flag"

	"github.com/hebl/gox/xweb"
	"golang.org/x/net/context"
)

//MyRouter simple router
type MyRouter struct {
	routes []web.Route
}

//Routes simple interface
func (mr *MyRouter) Routes() []web.Route {
	return mr.routes
}

//NewMyRouter new router
func NewMyRouter() web.Router {
	r := &MyRouter{}

	r.initRoutes()

	return r
}

func (mr *MyRouter) initRoutes() {
	mr.routes = []web.Route{
		//GET
		web.NewGetRoute("/get", mr.getHandler),
		//...
	}
}

//getHandler simple get handler
func (mr *MyRouter) getHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Write([]byte("get"))

	return nil
}

func main() {
	configFile := flag.String("c", "config.json", "JSON Config file")
	flag.Parse()

	router := NewMyRouter()
	server := web.New(*configFile)
	server.Init(router)
	server.StartServer()
}
