package ui

import (
	"net"
	"net/http"
	"time"
	"html/template"
	"github.com/luisfsantos/thysis/model"
	"fmt"
	"encoding/json"
	"log"
)

type Configuration struct {
	Assets http.FileSystem
}

type home struct {
	Title, CdnReact, CdnReactDom, CdnBabelStandalone, CdnAxios string
}

func Start(configuration Configuration, m *model.Model, listener net.Listener)  {
	server := &http.Server{
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 16}

	http.HandleFunc("/", homeHandler)
	http.Handle("/register", createUserHandler(m))
	http.Handle("/admin", viewUserHandler(m))
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

func createUserHandler(m *model.Model) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var user model.User
		err := decoder.Decode(&user)
		if err != nil {
			log.Printf("Error creating user: %v\n", err)
			http.Error(w, "Couldn't create user.", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		userID, err := m.DB.CreateUser(user.Username, user.Email, user.Password)
		if err != nil {
			log.Printf("Error creating user: %v\n", err)
			http.Error(w, "Couldn't create user", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, string(userID))
	})
}

func viewUserHandler(m *model.Model) http.Handler  {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := m.DB.SelectAllUsers()
		if err != nil {
			log.Printf("Error veiwing users: %v\n", err)
			http.Error(w, "Couldn't show users", http.StatusBadRequest)
			return
		}
		js, err := json.Marshal(users)
		fmt.Fprintf(w, string(js))
	})
}