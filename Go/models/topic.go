package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Topic struct {
	gorm.Model
	Title string `json:"title" validate:"required"`
	Content string `json:"content"`
}

func (t *Topic) String() {
	fmt.Printf("Title = %s, Content = %s \n", t.Title, t.Content)
}