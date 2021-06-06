package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"todo_app/pkg/api"
	"todo_app/pkg/todo"
)

func main()  {
	s := grpc.NewServer()
	srv := &todo.GRPCServer{}
	api.RegisterToDoServer(s,srv)

	todo.NewTask("Task 1. ", "Do your homework. ")
	todo.NewTask("Task 2. ", "Clean your room. ")
	todo.NewTask("Task 3. ", "Rest a little. ")
	todo.NewTask("Task 4. ", "Go to sleep. ")

	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}