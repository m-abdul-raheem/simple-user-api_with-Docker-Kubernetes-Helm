package main

type User struct {
	Id   int    `bson:"id,omitempty" validate:"required"`
	Name string `bson:"name,omitempty" validate:"required"`
}
