package handler

import (
	"encoding/json"
	"net/http"

	"github.com/seggga/backend2/internal/entity"
	"github.com/seggga/backend2/internal/logic/repo"
)

type Router struct {
	*http.ServeMux
	stor *repo.Storage
}

func NewRouter(stor *repo.Storage) *Router {
	r := &Router{
		ServeMux: http.NewServeMux(),
		stor:     stor,
	}

	r.Handle("/create-user", http.HandlerFunc(r.CreateUser))
	r.Handle("/create-group", http.HandlerFunc(r.CreateGroup))
	r.Handle("/add-to-group", http.HandlerFunc(r.AddToGroup))
	r.Handle("/remove-from-group", http.HandlerFunc(r.RemoveFromGroup))
	r.Handle("/search-user", http.HandlerFunc(r.SearchUser))
	r.Handle("/search-group", http.HandlerFunc(r.SearchGroup))

	return r
}

// CreateUser adds a new user, passed with post-request
func (rt *Router) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	u := entity.User{}
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newUser, err := rt.stor.Repo.CreateUser(r.Context(), u)
	if err != nil {
		http.Error(w, "error creating new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		entity.User{
			ID:    newUser.ID,
			Name:  newUser.Name,
			Email: newUser.Email,
		},
	)
}
func (rt *Router) CreateGroup(w http.ResponseWriter, r *http.Request)
func (rt *Router) AddToGroup(w http.ResponseWriter, r *http.Request)
func (rt *Router) RemoveFromGroup(w http.ResponseWriter, r *http.Request)
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request)
func (rt *Router) SearchGroup(w http.ResponseWriter, r *http.Request)
