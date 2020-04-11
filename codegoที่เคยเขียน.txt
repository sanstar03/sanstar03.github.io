package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	userDB := map[string]int{
		"Nissan": 22,
		"Golf":   23,
		"Benz":   24,
	}
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "signup.html")
	})
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "login.html")
	})
	router.HandleFunc("/File", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "name.txt")
	})

	router.HandleFunc("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := userDB[name]
		fmt.Fprintf(w, "My name is %s %d years old", name, age)
	}).Methods("GET")

	http.ListenAndServe(":8080", router)

}

/*func product(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "product request")
}*/
