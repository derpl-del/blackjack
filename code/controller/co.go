package controller

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/blackjack/code/page"
	"github.com/gorilla/mux"
)

var r *mux.Router

//CallController for generate api
func CallController() {
	fmt.Println("morning")
	r = mux.NewRouter()
	r.HandleFunc("/play", page.HelloWeb)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
