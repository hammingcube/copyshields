package main

import (
	"log"
	"net/http"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/go-martini/martini"
)


var last = "-correct-brightgreen.svg"

func viewHandler(w http.ResponseWriter, r *http.Request, params martini.Params) {
	if last == "-correct-brightgreen.svg" {
		last = "-wrong-red.svg"
	} else {
		last = "-correct-brightgreen.svg"
	}
	log.Println("redirecting")
	w.Header().Set("Cache-Control", "no-cache")
	http.Redirect(w, r, "https://img.shields.io/badge/solution"+last, 302)
}

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/", viewHandler)
	m.RunOnAddr(":8080")
	//m.Run(":8080")
}
