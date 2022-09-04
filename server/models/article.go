package models

type Article struct {
	Model
	Title   string `json: "title"`
	Desc    string `json: "desc"`
	Content string `json: "content"`
	State   string `json: "state"`
	Author  string `json: "author"`
}
