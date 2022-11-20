package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var languages = map[string]string{
	"eng": "Hello",
	"tur": "Merhaba",
}

type User struct {
	Name    string `json:"name,omitempty"`
	Surname string `json:"surname,omitempty"`
}

var LangNotExists = errors.New("language not exists")

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		lang := r.URL.Query().Get("language")

		resp, err := http.Get("http://user-service:8081/users?id=" + id)

		if err != nil {
			log.Println(err)
			return
		}

		if resp.StatusCode != http.StatusOK {
			log.Printf("HTTP status code %d\n", resp.StatusCode)
			return
		}

		body, _ := io.ReadAll(resp.Body)

		var user User
		err = json.Unmarshal(body, &user)

		if err != nil {
			log.Fatal(err)
		}

		gmsg, ok := languages[lang]

		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(LangNotExists.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(greet(user, gmsg)))

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(user User, greeting string) string {
	return fmt.Sprintf("%s, %s %s", greeting, user.Name, user.Surname)
}
