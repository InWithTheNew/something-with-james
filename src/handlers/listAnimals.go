package handlers

import (
	"animals/animals"
	"fmt"
	"net/http"
)

// This is the listening code from /animals/{animal} in main.go
func ListAnimals(w http.ResponseWriter, r *http.Request) {
	// Get the path and turn it into string.
	animalValue := r.PathValue("animal")

	// Create a new animal. The downstream code figures it out.
	animal := animals.New(animalValue)
	// If there is no animal in the 'animals' section, return error.
	if animal == nil {
		http.Error(w, "Animal not found", http.StatusNotFound)
		return
	}
	// Now we get the breeds and pass it back to the client.
	for _, breed := range animal.GetBreeds() {
		fmt.Fprintln(w, breed)
	}
}
