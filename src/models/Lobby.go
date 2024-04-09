package models

type Lobby struct {
	ID        string  `json:"id"`
	LobbyURL  string  `json:"lobby_url"`
	VideoURL  string  `json:"video_url"`
	CreatedAt string  `json:"created_at"`
	ChangedAt string  `json:"changed_at"`
	UserList  []*User `json:"user_list"`
}
