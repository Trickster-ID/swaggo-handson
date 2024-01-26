package model

type User struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	IsMarriage bool   `json:"is_marriage"`
	Title      string `json:"title"`
}
