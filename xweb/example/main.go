package main

import (
	"net/http"

	"flag"

	log "github.com/Sirupsen/logrus"
	"github.com/hebl/gox/xweb"
	"golang.org/x/net/context"
)

//
//PerType 是一个权限类型.
// 参考 https://github.com/jimmykuu/gopher/blob/master/permission.go
type PerType uint

const (
	// Everyone         访客             001
	Everyone PerType = 1 << 0

	// Authenticated    已登陆           010
	Authenticated = 1 << 1

	// AdministratorOnly 管理员           100
	AdministratorOnly = 1 << 2

	// Administrator    管理员也要已登陆   110
	Administrator = AdministratorOnly | Authenticated
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

func filterHandler(perm PerType, handler xweb.HTTPAPIFunc) xweb.HTTPAPIFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
		log.Infof("PerType: %v", perm)
		return handler(ctx, w, r, vars)
	}
}

func (mr *MyRouter) initRoutes() {
	mr.routes = []xweb.Route{
		//GET
		xweb.NewGetRoute("/get", filterHandler(Everyone, mr.getHandler)),
		//...
	}
}

//getHandler simple get handler
func (mr *MyRouter) getHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	w.Write([]byte("get"))

	return nil
}

func main() {
	flag.Parse()
	configFile := flag.String("c", "config.json", "JSON Config file")

	config := xweb.NewConfig(*configFile)
	router := NewMyRouter()
	server := xweb.NewServer(config, router)
	server.Init(router)
	server.StartServer()
}
