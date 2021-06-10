package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserModel represent a mgo database session with a user model data.
type UserModel struct {
	C *mongo.Collection
}

// All method will be used to get all records from the users table.
func (m *UserModel) All() ([]User, error) {
	// Define variables
	ctx := context.TODO()
	uu := []User{}

	// Find all users
	userCursor, err := m.C.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = userCursor.All(ctx, &uu)
	if err != nil {
		return nil, err
	}

	return uu, err
}

// FindByID will be used to find a new user registry by id
func (m *UserModel) FindByID(id int) (*User, error) {
	//p, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	return nil, err
	//}

	// Find user by id
	var user = User{}
	err := m.C.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		// Checks if the user was not found
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("ErrNoDocuments")
		}
		return nil, err
	}

	return &user, nil
}

// Insert will be used to insert a new user
func (m *UserModel) Insert(user User) (*mongo.InsertOneResult, error) {
	return m.C.InsertOne(context.TODO(), user)
}
