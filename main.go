package main

import (
	"fmt"
	"goc/logface"
	"goc/toolcom/cfgtool"
	"goc/toolcom/errtool"
	"goc/toolcom/jsontool"
	"net/http"
)

var log = logface.New(logface.DebugLevel)

var serverCfg struct {
	Port int
	Path string
}

func main() {
	cfg := cfgtool.New("conf.json")
	json, err := cfg.TakeJson("SerCfg")
	errtool.Errpanic(err)
	err = jsontool.Decode(json, &serverCfg)
	errtool.Errpanic(err)
	log.Info("file server running[%+v]", serverCfg)
	http.ListenAndServe(fmt.Sprintf(":%d", serverCfg.Port), http.FileServer(http.Dir(serverCfg.Path)))

}
