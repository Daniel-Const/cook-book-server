package main

import "fmt"

func main() {
    
    recipe := &Recipe{"Ravioli", "Tasty classic", []string{"Pasta", "Sauce"}};
    recipe.Save();
    var r Recipe
    err := LoadRecipe("Ravioli.txt", &r);
    if (err != nil) {
        fmt.Println("Failed to load recipe!");
    }
    fmt.Println(r.Title);
}
