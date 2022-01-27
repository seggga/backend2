package handler

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/seggga/backend2/internal/entity"
	"github.com/seggga/backend2/internal/logic/repo"
)

type Router struct {
	*http.ServeMux
	stor *repo.Storage
}

// NewRouter creates a router with specified storage and handlers
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
		http.Error(w, "method mismatch", http.StatusMethodNotAllowed)
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
		http.Error(w, "method mismatch", http.StatusMethodNotAllowed)
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

// AddToGroup add user specified into the group, passed with get-request
// .../add-to-group?uid=...&gid=...
func (rt *Router) AddToGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method mismatch", http.StatusMethodNotAllowed)
		return
	}

	// check and parse parameters
	uidParameter := r.URL.Query().Get("uid")
	if uidParameter == "" {
		http.Error(w, "uid should be set", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(uidParameter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	gidParameter := r.URL.Query().Get("gid")
	if gidParameter == "" {
		http.Error(w, "gid should be set", http.StatusBadRequest)
		return
	}
	gid, err := uuid.Parse(gidParameter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (gid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = rt.stor.Repo.AddToGroup(r.Context(), uid, gid)
	if err != nil {
		http.Error(w, "error adding user to the group", http.StatusInternalServerError)
		return
	}

}

// RemoveFromGroup removes user specified from the group, passed with get-request
// .../remove-from-group?uid=...&gid=...
func (rt *Router) RemoveFromGroup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method mismatch", http.StatusMethodNotAllowed)
		return
	}

	// check and parse parameters
	uidParameter := r.URL.Query().Get("uid")
	if uidParameter == "" {
		http.Error(w, "uid should be set", http.StatusBadRequest)
		return
	}
	uid, err := uuid.Parse(uidParameter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (uid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	gidParameter := r.URL.Query().Get("gid")
	if gidParameter == "" {
		http.Error(w, "gid should be set", http.StatusBadRequest)
		return
	}
	gid, err := uuid.Parse(gidParameter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (gid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = rt.stor.Repo.RemoveFromGroup(r.Context(), uid, gid)
	if err != nil {
		http.Error(w, "error adding user to the group", http.StatusInternalServerError)
		return
	}
}

// SearchUser searches users by name or by group specified, passed with get-request
// .../search-user?name=...&gid1=...&gid2=...&gid3=...
func (rt *Router) SearchUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method mismatch", http.StatusMethodNotAllowed)
		return
	}

	// check and parse parameters
	name := r.URL.Query().Get("name")

	gids := make([]uuid.UUID)
	gidParams := make([]string)
	gidParameter := r.URL.Query().Get("gid1")
	if gidParameter != "" {
		gidParams = append(gidParams, gidParameter)
	}

	gidParameter = r.URL.Query().Get("gid2")
	if gidParameter != "" {
		gidParams = append(gidParams, gidParameter)
	}

	gidParameter = r.URL.Query().Get("gid3")
	if gidParameter != "" {
		gidParams = append(gidParams, gidParameter)
	}

	/// .... проверить, что передано в параметрах
	gidParameter := r.URL.Query().Get("gid")
	if gidParameter == "" {
		http.Error(w, "gid should be set", http.StatusBadRequest)
		return
	}
	gid, err := uuid.Parse(gidParameter)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if (gid == uuid.UUID{}) {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

}

func (rt *Router) SearchGroup(w http.ResponseWriter, r *http.Request)
