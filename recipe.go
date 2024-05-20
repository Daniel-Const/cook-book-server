package main

import (
    "os"
    "encoding/json"
)

type Recipe struct {
    Title string
    Description string
    Ingredients []string
}

func (r *Recipe) Serialize() ([]byte, error) {
    return json.Marshal(r);
}

func (r *Recipe) Save() error {
    filename := r.Title + ".txt"
    bytes, _  := r.Serialize();
    return os.WriteFile(filename, bytes, 0600);
}

func LoadRecipe(filename string, recipe *Recipe) error {
    bytes, _ := os.ReadFile(filename);
    err := json.Unmarshal(bytes, &recipe);
    return err
}
