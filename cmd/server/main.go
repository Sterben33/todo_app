package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"todo_app/pkg/api"
	"todo_app/pkg/todo"
)

func main()  {
	//Creating server
	s := grpc.NewServer()
	srv := &todo.GRPCServer{}
	api.RegisterToDoServer(s,srv)

	// Stop server when interrupted
	defer s.Stop()

	// Some tasks that will be available at the very beginning
	todo.NewTask("Task 1. ", "Do your homework. ")
	todo.NewTask("Task 2. ", "Clean your room. ")
	todo.NewTask("Task 3. ", "Rest a little. ")
	todo.NewTask("Task 4. ", "Go to sleep. ")

	l, err := net.Listen("tcp", ":5050")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}