package  models

import (
	"fmt"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name string `json:"name" form:"name" validate:"required"`
}

func (t *Category) String() {
	fmt.Printf("Name = %s \n", t.Name)
}