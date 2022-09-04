package models

type ContentTag struct {
	Model
	ContentId string `json: "content_id"`
	TagId     string `json: "tag_id"`
}
