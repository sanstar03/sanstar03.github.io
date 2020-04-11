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
	router.HandleFunc("/move", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "scan.html")
	})
	router.HandleFunc("/scan/stop", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "passwordreq.html")
	})
	router.HandleFunc("/scan/stop/code", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "correctpass.html")
	})

	router.HandleFunc("/scan/stop/code/open", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "openrobot.html")
	})
	router.HandleFunc("/scan/stop/code/open/close", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robotready.html")
	})
	router.HandleFunc("/scan/stop/code/open/close/move", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robotmoveafterclose.html")
	})

	router.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "fistcheck.html")
	})
	router.HandleFunc("/show", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "firstshow.html")
	})
	router.HandleFunc("/secondcheck", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "secondcheck.html")
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
