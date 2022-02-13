package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/seggga/backend2/internal/red"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB

	measurable = red.MeasurableHandler

	router = mux.NewRouter()
	web    = http.Server{
		Handler: router,
	}
)

func main() {
	router.
		HandleFunc("/entities", measurable(ListEntitiesHandler)).
		Methods(http.MethodGet)
	router.
		HandleFunc("/entity", measurable(AddEntityHandler)).
		Methods(http.MethodPost)
	var err error
	db, err = sql.Open("mysql", "root:test@tcp(mysql:3306)/test")
	if err != nil {
		panic(err)
	}

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(":9090", nil); err != http.ErrServerClosed {
			panic(fmt.Errorf("error on listen and serve: %v", err))
		}
	}()
	if err := web.ListenAndServe(); err != http.ErrServerClosed {
		panic(fmt.Errorf("error on listen and serve: %v", err))
	}
}

const sqlInsertEntity = `
  INSERT INTO entities(id, data) VALUES (?, ?)
  `

// AddEntityHandler ...
func AddEntityHandler(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(fmt.Sprintf("http://acl/identity?token=%s", r.FormValue("token")))
	switch {
	case err != nil:
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	case res.StatusCode != http.StatusOK:
		w.WriteHeader(http.StatusForbidden)
		return
	}
	res.Body.Close()

	_, err = db.Exec(sqlInsertEntity, r.FormValue("id"), r.FormValue("data"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

const sqlSelectEntities = `
  SELECT id, data FROM entities;
  `

// ListEntityItemResponse ...
type ListEntityItemResponse struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

// ListEntitiesHandler ...
func ListEntitiesHandler(w http.ResponseWriter, r *http.Request) {
	rr, err := db.Query(sqlSelectEntities)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer rr.Close()

	var ii = []*ListEntityItemResponse{}
	for rr.Next() {
		i := &ListEntityItemResponse{}
		err = rr.Scan(&i.ID, &i.Data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		ii = append(ii, i)
	}
	bb, err := json.Marshal(ii)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(bb)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
