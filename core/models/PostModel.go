package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title string
	Content string
}

type ApiPost struct {
	ID uint `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
}