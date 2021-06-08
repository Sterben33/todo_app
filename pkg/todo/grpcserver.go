package todo

import (
	"context"
	"errors"
	"todo_app/pkg/api"
)

// GRPCServer ...
type GRPCServer struct {}

// TaskMap Here is our "DB"
var TaskMap = make(map[int32]api.Task)

// Counter Auto increment
var Counter int32 = 0

// NewTask constructor
func NewTask(Title,Body string) api.Task {
	defer func() {Counter++}() // Auto increment id
	TaskMap[Counter] = api.Task{Id: Counter, Title:Title, Body: Body, Done: false}
	return TaskMap[Counter]
}

func (s *GRPCServer) Create(_ context.Context, request *api.Data) (*api.Task, error) {
	t := NewTask(request.Title, request.Body)
	return &t, nil
}

func (s *GRPCServer) Read(_ context.Context, request *api.TaskId) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		return &t, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) Update(_ context.Context, request *api.Task) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		t.Title = request.GetTitle()
		t.Body = request.GetBody()
		TaskMap[request.GetId()] = t
		return &t, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) Delete(_ context.Context, request *api.TaskId) (*api.Empty, error) {
	if _, exist := TaskMap[request.GetId()]; exist {
		delete(TaskMap, request.GetId())
		return &api.Empty{}, nil
	}
	return &api.Empty{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) MarkAsDone(_ context.Context, request *api.TaskId) (*api.Task, error) {
	if t, exist := TaskMap[request.GetId()]; exist {
		t.Done = true
		TaskMap[request.GetId()] = t
		return &t, nil
	}
	return &api.Task{}, errors.New("Task with such ID does not exist. ")
}

func (s *GRPCServer) GetList(_ context.Context, _ *api.Empty) (*api.TaskList, error) {
	l := make([]*api.Task,0,len(TaskMap))
	for i, _ := range TaskMap {
		t := TaskMap[i]
		l = append(l, &t)
	}
	return &api.TaskList{Task: l}, nil
}
