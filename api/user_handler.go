package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/enkelm/go_api/db"
	"github.com/enkelm/go_api/db/model"

	"github.com/gorilla/mux"
)

const base = "/users"

func addUserHandlers(r *router) {
	r.handleGet(CreateUrl(base), getUsers)
	r.handleGet(CreateUrl(base, "{id}"), getUser)
	r.handleGet(CreateUrl(base, "signup"), signUp)
	r.handleGet(CreateUrl(base, "signin", "{email}", "{pass}"), signIn)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	pg, _ := db.NewPGInstance(context.Background())

	user, _ := pg.GetUsers(context.Background())
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	pg, _ := db.NewPGInstance(context.Background())
	vars := mux.Vars(r)

	user, err := pg.GetUserById(vars["id"])
	if err != nil {
		log.Println("Couldn't retrieve user: %w", err)
	}
	defer pg.Close()

	json.NewEncoder(w).Encode(user.Data())
}

func signUp(w http.ResponseWriter, r *http.Request) {
	pg, _ := db.NewPGInstance(context.Background())

	err := pg.SignUp(context.Background(), &model.UserDto{
		FirstName: "Test",
		LastName:  "Test",
		Email:     "test@test.com",
	}, "P@ssword1")

	json.NewEncoder(w).Encode(err == nil)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	pg, _ := db.NewPGInstance(context.Background())
	vars := mux.Vars(r)

	isValid := pg.SignIn(context.Background(), vars["email"], vars["pass"])

	json.NewEncoder(w).Encode(isValid)
}
