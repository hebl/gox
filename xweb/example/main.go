package main

import (
	"net/http"

	"flag"

	"github.com/hebl/gox/xweb"
	"golang.org/x/net/context"
)

//MyRouter simple router
type MyRouter struct {
	routes []xweb.Route
}

//Routes simple interface
func (mr *MyRouter) Routes() []xweb.Route {
	return mr.routes
}

//NewMyRouter new router
func NewMyRouter() xweb.Router {
	r := &MyRouter{}

	r.initRoutes()

	return r
}

func (mr *MyRouter) initRoutes() {
	mr.routes = []xweb.Route{
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
	server := xweb.New(*configFile)
	server.Init(router)
	server.StartServer()
}
