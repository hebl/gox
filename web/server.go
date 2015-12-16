package web

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/context"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/hebl/gox/utils"
)

//Server Web程序
type Server struct {
	config *Config
	router Router
}

//New 新建一个Web实例
func New(configFile string) *Server {
	var config Config
	utils.ParseJSONFile(configFile, &config)

	server := &Server{
		config: &config,
	}
	server.init()
	return server
}

//Init 初始化 router
func (s *Server) Init(router Router) {
	s.router = router
}

func (s *Server) init() {
	//Log
	log.SetFormatter(&log.TextFormatter{})

	// Log Output
	if s.config.LogFile != "" {
		logFile, err := os.Open(s.config.LogFile)
		if err != nil {
			log.Errorf("Open log file %s error : %v", s.config.LogFile, err)
		}
		log.SetOutput(logFile)
	} else {
		log.SetOutput(os.Stderr)
	}

	log.SetLevel(s.config.LogLevel)

}

// CreateMux 根据router生成http mux
func (s *Server) CreateMux() *mux.Router {
	m := mux.NewRouter()

	for _, r := range s.router.Routes() {

		f := makeHTTPHandler(r.Handler)

		log.Debugf("Registering %s, %s", r.Method, r.Path)
		m.Path(r.Path).Methods(r.Method).HandlerFunc(f)

	}

	//绑定静态目录
	if s.config.Static != "" {
		log.Debugf("Registering Static directory %s", s.config.Static)
		m.PathPrefix(s.config.Static).HandlerFunc(FileHandler)

	}

	return m
}

func makeHTTPHandler(handler HTTPAPIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Calling %s %s", r.Method, r.URL.Path)

		ctx := context.Background()
		handlerFunc := handler

		if err := handlerFunc(ctx, w, r, mux.Vars(r)); err != nil {
			log.Errorf("Handler for %s %s returned error: %v", r.Method, r.URL.Path, err)
			WriteError(w, err)
		}
	}
}

//StartServer 启动服务
func (s *Server) StartServer() {

	http.Handle("/", s.CreateMux())
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.config.Port), nil)
	if err != nil {
		log.Errorf("服务器启动错误： %v", err)
	}
	log.Debugf("服务器启动 :%d", s.config.Port)
}
