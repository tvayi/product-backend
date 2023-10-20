package models

import "github.com/go-pg/pg/v10"

type Task struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    Duration int64 `json:"duration"`
}

func CreateTask(db *pg.DB, req *Task) (*Task, error) {
    _, err := db.Model(req).Insert()
    if err != nil {
        return nil, err
    }

    task := &Task{}

    err = db.Model(task).
        Where("task.id = ?", req.ID).
        Select()

    return task, err
}

func GetTask(db *pg.DB, taskID string) (*Task, error) {
    task := &Task{}

    err := db.Model(task).
        Where("task.id = ?", taskID).
        Select()

    return task, err
}

func GetTasks(db *pg.DB) ([]*Task, error) {
    tasks := make([]*Task, 0)

    err := db.Model(&tasks).
        Select()

    return tasks, err
}
