package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

const (
	DefaultPort = "8080"
	RecipeRoute = "/recipe/"
)

// /recipe/ Handler
func recipeHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case http.MethodGet:
		encoder := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json")

		recipeName := r.URL.Path[len(RecipeRoute):]

		// Fetch all recipes
		if recipeName == "" {
			recipes := LoadAllRecipes()
			encoder.Encode(recipes)
			return
		}

		log.Println("Fetching: ", recipeName)
		recipe, err := LoadRecipe(recipeName + ".txt")
		if err != nil {
			log.Println("Failed to load Recipe")
			fmt.Fprintf(w, "Failed to load recipe: %s", err.Error())
			return
		}
		encoder.Encode(recipe)

	case http.MethodPost:
		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var recipe Recipe
		err := d.Decode(&recipe)
		if err != nil {
			http.Error(w, "Oopsie", http.StatusBadRequest)
			return
		}
		recipe.Save()

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc(RecipeRoute, recipeHandler)
	log.Printf("Server is litening on port %s...", DefaultPort)
	log.Fatal(http.ListenAndServe(":"+DefaultPort, nil))
}
