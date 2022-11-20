package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
)

type User struct {
	Id      int    `json:"id,omitempty"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type DB map[int]User

func (d DB) Get(id int) (*User, error) {
	rec, ok := d[id]

	if !ok {
		return nil, errors.New("record already exists")
	}

	return &rec, nil
}

func main() {
	db := DB{}
	fillDB(db)

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		u, err := db.Get(id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))

		}

		m, _ := json.Marshal(u)
		w.WriteHeader(http.StatusOK)
		w.Write(m)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func fillDB(db DB) {
	users := []User{
		{Id: 1, Name: "Rob", Surname: "Pike"},
		{Id: 2, Name: "Robert", Surname: "Griesemer"},
	}

	for _, u := range users {
		db[u.Id] = u
	}
}
