package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
	FeatureImageURL string `json:"featureImageURL" validate:"required"`
}

func (t *Blog) String() {
	fmt.Printf("Title = %s, Content = %s, Feature Image URL = %s \n", t.Title, t.Content, t.FeatureImageURL)
}