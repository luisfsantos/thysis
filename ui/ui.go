package ui

import (
	"net"
	"net/http"
	"time"
	"html/template"
)

type Configuration struct {
	Assets http.FileSystem
}

type home struct {
	Title, CdnReact, CdnReactDom, CdnBabelStandalone, CdnAxios string
}

func Start(configuration Configuration, listener net.Listener)  {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.HandleFunc("/", homeHandler)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(configuration.Assets)))

	go server.Serve(listener)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	p := home{
		Title:"Home Page",
		CdnReact:"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js",
		CdnReactDom:"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js",
		CdnBabelStandalone:"https://cdnjs.cloudflare.com/ajax/libs/babel-standalone/6.24.0/babel.min.js",
		CdnAxios:"https://cdnjs.cloudflare.com/ajax/libs/axios/0.16.1/axios.min.js"}
	t, _ := template.ParseFiles("grayscale/index.html")
	t.Execute(w, p)
}