package repositories

import (
	"context"
	"errors"
	domain "test_task_manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (t *taskRepository) CreateTask(c context.Context, newTask domain.Task) (*domain.Task, error) {
	id := newTask.ID
	_, err := t.GetTaskByID(c, id)
	if err == nil {
		return nil, errors.New("task with the given id already exists")
	}

	collection := t.database.Collection(t.collection)

	_, err = collection.InsertOne(c, newTask)

	if err != nil {
		return nil, err
	}
	return &newTask, nil
}

func (t *taskRepository) DeleteTask(c context.Context, taskID string) error {
	collection := t.database.Collection(t.collection)

	filter := bson.D{{Key: "id", Value: taskID}}
	_, err := collection.DeleteOne(c, filter)
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetTaskByID(c context.Context, taskID string) (*domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var task domain.Task

	filter := bson.D{{Key: "id", Value: taskID}}
	err := collection.FindOne(c, filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetTasks(c context.Context) ([]domain.Task, error) {
	collection := t.database.Collection(t.collection)

	var tasks []domain.Task

	cur, err := collection.Find(c, bson.D{{}})

	if err != nil {
		return nil, err
	}

	for cur.Next(c) {
		var task domain.Task
		if err := cur.Decode(&task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	cur.Close(c)

	return tasks, nil
}

func (t *taskRepository) UpdateTask(c context.Context, taskID string, updatedTask domain.Task) (*domain.Task, error) {
	collection := t.database.Collection(t.collection)

	filter := bson.D{{Key: "id", Value: taskID}}
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

	result := collection.FindOneAndUpdate(context.TODO(), filter, update, options.FindOneAndUpdate().SetReturnDocument(options.After))
	if result.Err() != nil {
		return nil, result.Err()
	}

	var taskAfterUpdate domain.Task
	err := result.Decode(&taskAfterUpdate)
	if err != nil {
		return nil, err
	}

	return &taskAfterUpdate, nil
}
