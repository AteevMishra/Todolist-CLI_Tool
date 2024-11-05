package db

import (
	"context"
	"errors"

	"github.com/AteevMishra/todo-CLI/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAllTodos retrieves all todos from the collection
func GetAllTodos() ([]models.Todo, error) {
	var todos []models.Todo
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

// AddTodo adds a new todo to the collection
func AddTodo(todo *models.Todo) error {
	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)
	return nil
}

// UpdateTodoStatus updates a todo's completion status
func UpdateTodoStatus(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	return err
}

// DeleteTodoByID deletes a todo by its ID
func DeleteTodoByID(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	filter := bson.M{"_id": objectId}
	_, err = collection.DeleteOne(context.Background(), filter)
	return err
}