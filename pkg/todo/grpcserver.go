package todo

import (
	"context"
	"errors"
	"todo_app/pkg/api"
)

type GRPCServer struct {}

type Task struct {
	id int32
	title string
	body string
	done bool
}

var TaskList = make(map[int32]Task)

var Counter int32 = 0

func NewTask(Title,Body string) Task {
	defer func() {Counter++}()
	TaskList[Counter] = Task{id:Counter, title:Title, body: Body, done: false}
	return TaskList[Counter]
}


func (s *GRPCServer) Get(ctx context.Context, request *api.GetRequest) (*api.GetResponce, error) {
	if t, err := TaskList[request.GetId()]; err == true {
		return &api.GetResponce{Id: t.id, Title: t.title, Body: t.body, Done: t.done}, nil
	}
	return &api.GetResponce{}, errors.New("Task with such ID does not exist. ")
}
