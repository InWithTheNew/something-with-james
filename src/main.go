package main

import (
	"animals/handlers"
	"net/http"
)

// This kicks off the code.
func main() {
	// This says we're starting a new router.
	r := http.NewServeMux()
	// This is an 'endpoint' or 'route'. These define what it's listening on and where it'll route it.
	// So in this case http://localhost:8080/animals.
	r.HandleFunc("GET /animals/{animal}", handlers.ListAnimals)
	// This starts the server and tells us what port to use.
	http.ListenAndServe(":8080", r)
}
