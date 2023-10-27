package model

type TodoDetail struct {
	ID     string `json:"id,omitempty" :"id"`
	Text   string `json:"text" json:"text,omitempty" :"text"`
	UserID string `json:"userId" json:"user_id,omitempty" :"user_id"`
}
