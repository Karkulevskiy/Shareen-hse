package domain

type User struct {
	ID       int64  `json:"id,omitempty"`
	Login    string `json:"login,omitempty"`
	PassHash []byte `json:"pass_hash,omitempty"`
}
