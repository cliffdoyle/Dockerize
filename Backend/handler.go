package backend

import (
	"net/http"
	"text/template"
)

func Webhandler(write http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(write, "404 not found", http.StatusNotFound)
		return
	}
	temp := template.Must(template.ParseFiles("Templates/index.html"))
	temp.Execute(write, nil)
}
