package main

import (
	"fmt"
	"log"
	"net/http"
    "path"
    "encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {

    switch r.Method {
    case http.MethodGet:
        recipeName := path.Base(r.URL.Path)
        log.Println("Fetching: ", recipeName)
        recipe, err := LoadRecipe(recipeName + ".txt")
        if err != nil {
            log.Println("Failed to load Recipe")
            fmt.Fprintf(w, "Failed to load recipe: %s", err.Error())
            return
        }
        fmt.Fprintf(w, "Title: %s\nDescription: %s\nIngredients: %v", recipe.Title, recipe.Description, recipe.Ingredients)
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
    http.HandleFunc("/recipe/", recipeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
