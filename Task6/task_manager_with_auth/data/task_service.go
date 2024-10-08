package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetTasks() ([]models.Task, error) {
	var tasks []models.Task
	cursor, err := taskCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var task models.Task
		if err := cursor.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(context.TODO())

	return tasks, nil
}

func GetTaskByID(id string) (*models.Task, error) {
	var task models.Task

	filter := bson.M{"id": id}
	err := taskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func CreateTask(newTask models.Task) (*models.Task, error) {
	id := newTask.ID
	_, err := GetTaskByID(id)
	if err == nil {
		return nil, errors.New("task with the given id already exists")
	}
	_, err = taskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func UpdateTask(id string, updatedTask models.Task) (*models.Task, error) {
	filter := bson.M{"id": id}

	updateFields := bson.D{}
	if updatedTask.Title != "" {
		updateFields = append(updateFields, bson.E{Key: "title", Value: updatedTask.Title})
	}
	if updatedTask.Description != "" {
		updateFields = append(updateFields, bson.E{Key: "description", Value: updatedTask.Description})
	}
	if updatedTask.Status != "" {
		updateFields = append(updateFields, bson.E{Key: "status", Value: updatedTask.Status})
	}
	if !updatedTask.DueDate.IsZero() {
		updateFields = append(updateFields, bson.E{Key: "due_date", Value: updatedTask.DueDate})
	}

	update := bson.D{{Key: "$set", Value: updateFields}}

	result := taskCollection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var updated models.Task
	err := result.Decode(&updated)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeleteTask(id string) error {
	filter := bson.M{"id": id}
	_, err := taskCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}
