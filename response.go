package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/yaoshifeng/gowebvideoserver/tree/master/api/defs"
)

//给客户端返回对应错误code
func sendErrorResponse(w http.ResponseWriter, errRep def.ErrorReponse) {
	w.WriteHeader(errRep.HttpSC)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	status, _ := json.Marshal(&errRep.Error)
	io.WriteString(w, string(status))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
