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

var TaskMap = make(map[int32]Task)

var Counter int32 = 0

func NewTask(Title,Body string) Task {
	defer func() {Counter++}()
	TaskMap[Counter] = Task{id: Counter, title:Title, body: Body, done: false}
	return TaskMap[Counter]
}

func (s *GRPCServer) Create(ctx context.Context, request *api.Data) (*api.Task, error) {
	t := NewTask(request.Title, request.Body)
	return &api.Task{Id: t.id, Title: t.title, Body: t.body, Done: t.done}, nil
}

func (s *GRPCServer) Read(ctx context.Context, request *api.TaskId) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		return &api.Task{Id: t.id, Title: t.title, Body: t.body, Done: t.done}, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) Update(ctx context.Context, request *api.Task) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		TaskMap[request.GetId()] = Task{id: t.id, title: request.GetTitle(), body: request.GetBody(), done: t.done}
		return &api.Task{Id: t.id, Title: request.GetTitle(), Body: request.GetBody(), Done: t.done}, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) Delete(ctx context.Context, request *api.TaskId) (*api.Empty, error) {
	if _, exist := TaskMap[request.GetId()]; exist {
		delete(TaskMap, request.GetId())
		return &api.Empty{}, nil
	}
	return &api.Empty{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) MarkAsDone(ctx context.Context, request *api.TaskId) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		t.done = true
		TaskMap[request.GetId()] = t
		return &api.Task{Id: t.id, Title: t.title, Body: t.body, Done: t.done}, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) GetList(ctx context.Context, request *api.Empty) (*api.TaskList, error) {
	l := make([]*api.Task,len(TaskMap))
	for _, task := range TaskMap {
		l = append(l, &api.Task{Id: task.id, Title: task.title, Body: task.body, Done: task.done})
	}
	return &api.TaskList{Task: l}, nil
}
