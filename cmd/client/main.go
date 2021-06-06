package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"strconv"
	"todo_app/pkg/api"
)

func main() {
	fmt.Println("Input id of the task. ")
	var inputstr string

	_, err := fmt.Scanf("%s", &inputstr)
	if err != nil {
		fmt.Println(err)
	}
	id, e := strconv.Atoi(inputstr)
	if e != nil {
		fmt.Println(e)
	}

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewToDoClient(conn)
	res, err := c.Get(context.Background(), &api.GetRequest{Id: int32(id)})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Re %s",res.GetBody())
}