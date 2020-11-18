package page

import (
	"fmt"
	"net/http"
)

//HelloWeb intro
func HelloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "text")
}
