// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CharacterResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}

type CreateCharacterRequest struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
