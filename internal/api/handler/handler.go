package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
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

// CreateGroup adds a new group, passed with post-request
func (rt *Router) CreateGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	g := entity.Group{}
	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	newGroup, err := rt.stor.Repo.CreateGroup(r.Context(), g)
	if err != nil {
		http.Error(w, "error creating new user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

	_ = json.NewEncoder(w).Encode(
		entity.Group{
			ID:          newGroup.ID,
			Name:        newGroup.Name,
			Description: newGroup.Description,
		},
	)
}

// AddToGroup add user specified into the group, passed wity get-request
// .../add-to-group?uid=...&gid=...
func (rt *Router) AddToGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "bad method", http.StatusMethodNotAllowed)
		return
	}

	//	...

	suid := r.URL.Query().Get("uid")
	if suid == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(suid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	nbu, err := rt.us.Read(r.Context(), uid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "not found", http.StatusNotFound)
		} else {
			http.Error(w, "error when reading", http.StatusInternalServerError)
		}
		return
	}

	_ = json.NewEncoder(w).Encode(
		User{
			ID:         nbu.ID,
			Name:       nbu.Name,
			Data:       nbu.Data,
			Permission: nbu.Permissions,
		},
	)
}

func (rt *Router) RemoveFromGroup(w http.ResponseWriter, r *http.Request)
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request)
func (rt *Router) SearchGroup(w http.ResponseWriter, r *http.Request)
