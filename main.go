package main

import (
	"fmt"
	"log"
	"net/http"
    "path"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func recipeHandler(w http.ResponseWriter, r *http.Request) {
    recipeName := path.Base(r.URL.Path)
    fmt.Println(recipeName)
    var recipe Recipe
    LoadRecipe(recipeName + ".txt", &recipe)
    fmt.Fprintf(w, "Title: %s\nDescription: %s\nIngredients: %v", recipe.Title, recipe.Description, recipe.Ingredients)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/recipe/", recipeHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
