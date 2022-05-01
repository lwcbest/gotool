package core

import (
	"github.com/lwcbest/gotool/que_str"
	"github.com/lwcbest/gotool/usejs"
	"log"
	"net/http"
)

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := que_str.ReqStr()

	//resp = "<table style='border:1px solid #F00; line-height: 25px; text-align: center; border-collapse: collapse;'><tr><td>指标</td></tr><tr><td>123</td><td>123</td></tr><tr><td>123</td><td>23123</td></tr></table>"
	w.Write([]byte(resp))
}

type httpServer2 struct {
}

func (server2 httpServer2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := que_str.SaveReqStr()
	w.Write([]byte(resp))
}

func StartReqStrServ() {
	usejs.UseJs()
	var server httpServer
	var server2 httpServer2
	http.Handle("/mohe", server)
	http.Handle("/mohe_save", server2)
	log.Println("http://localhost:8123/mohe")
	log.Println("http://localhost:8123/mohe_save")
	log.Fatal(http.ListenAndServe(":8123", nil))
}
