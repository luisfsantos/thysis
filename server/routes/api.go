package routes

import (
	"github.com/luisfsantos/thysis/model"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"log"
	"fmt"
)

func SetAPIRoutes(m *mux.Router, model *model.Model) *mux.Router {
	api := m.PathPrefix("/api").Subrouter()
	api.Path("/register").Methods("Post").Handler(createUserHandler(model))
	api.Path("/user/view/all").Methods("GET").Handler(viewUserHandler(model))
	return m
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
		err = m.DB.CreateUser(user.Username, user.Email, user.Password)
		if err != nil {
			log.Printf("Error creating user: %v\n", err)
			http.Error(w, "Couldn't create user", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User Created!")
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