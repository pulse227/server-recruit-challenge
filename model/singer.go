package model

type SingerID int

type Singer struct {
	ID   SingerID `json:"id"`
	Name string   `json:"name"`
}
