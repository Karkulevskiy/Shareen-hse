package models

type User struct {
	ID      string `json:"id"`
	LobbyID string `json:"lobby_id"`
	Name    string `json:"name"`
}
