package handler

import (
	"net/http"
	"io"
	"github.com/pineda89/golang-springboot-archetype/service"
	"strconv"
)

func StartWebServer(port int) {
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":" + strconv.Itoa(port), nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, service.MyApplicationMethod())
}
