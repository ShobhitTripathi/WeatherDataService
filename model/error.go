package model

type Error struct {
	Code uint   `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
	Type string `json:"type,omitempty"`
}
