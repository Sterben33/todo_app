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
	return Task{id:Counter, title:Title, body: Body, done: false}
}


func (s *GRPCServer) Get(ctx context.Context, request *api.GetRequest) (*api.GetResponce, error) {
	if t, err := TaskList[request.Id]; err == false {
		return &api.GetResponce{Id: t.id, Title: t.title, Body: t.body, Done: t.done}, nil
	}
	return &api.GetResponce{}, errors.New("Task with such ID does not exist. ")
}
