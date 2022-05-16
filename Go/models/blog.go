package models

import (
	"fmt"
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title string `json:"title" form:"title" validate:"required"`
	Content string `json:"content" form:"content" validate:"required"`
	FeatureImageURL string `json:"featureImageURL" form:"featureImageURL" validate:"required"`
	CategoryID int `json:"category_id" form:"category_id"`
	Category Category
}

func (t *Blog) String() {
	fmt.Printf("Title = %s, Content = %s, Feature Image URL = %s, Category ID = %d \n", t.Title, t.Content, t.FeatureImageURL, t.CategoryID)
}