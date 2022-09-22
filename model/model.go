package model

import "mime/multipart"

type Request struct {
	URL      string                `json:"url" form:"url"`
	ShortURL string                `json:"short_url" form:"short_url"`
	NewURL   string                `json:"new_url" form:"new_url"`
	Path     string                `json:"path" form:"path"`
	File     *multipart.FileHeader `json:"file" form:"file"`
}
type Reply struct {
	URL      string `json:"url"`
	ShortURL string `json:"short_url"`
}
