package main

import (
	"encoding/json"
	"log"
	"os"
)

type Recipe struct {
	Title       string
	Description string
	Ingredients []string
}

func (r *Recipe) Serialize() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Recipe) Save() error {
	filename := "data/" + r.Title + ".txt"
	bytes, err := r.Serialize()
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0600)
}

func LoadRecipe(filename string) (*Recipe, error) {
	bytes, err := os.ReadFile("data/" + filename)
	if err != nil {
		return nil, err
	}

	var recipe Recipe
	err = json.Unmarshal(bytes, &recipe)
	return &recipe, err
}

func LoadAllRecipes() []*Recipe {
	var recipes []*Recipe
	var recipe *Recipe

	files, err := os.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		recipe, err = LoadRecipe(file.Name())
		if err != nil {
			log.Fatal(err)
		}
		recipes = append(recipes, recipe)
	}

	return recipes
}
