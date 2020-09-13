package main

import "github.com/jinzhu/gorm"

// todoModel describes a todoModel type
type todoModel struct {
	gorm.Model
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// transformedTodo represents a formatted todo
// for response to the api
type transformedTodo struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
