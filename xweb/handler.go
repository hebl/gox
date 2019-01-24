package xweb

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

//FileHandler 静态文件
func FileHandler(w http.ResponseWriter, req *http.Request) {
	url := req.Method + " " + req.URL.Path
	log.Debugf("File URL: %s", url)
	filePath := req.URL.Path[1:]
	http.ServeFile(w, req, filePath)
}

// HandlerFunc HandlerFunc
type HandlerFunc func(*Context)

// HandlersChain HandlersChain
type HandlersChain []HandlerFunc
