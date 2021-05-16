package models

import "gorm.io/gorm"

type User struct{
    gorm.Model
    Email string `gorm:"unique_index;not null"`
    Name string  
}