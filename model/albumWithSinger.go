package model

type AlbumWithSinger struct {
	ID       AlbumID  `json:"id"`
	Title    string   `json:"title"`
	Singer   Singer   `json:"singer"` // モデル Singerと紐づきます
}