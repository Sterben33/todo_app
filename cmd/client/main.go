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

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewToDoClient(conn)

	for {
		fmt.Println("Choose an option:\n " +
			"list - get list of the tasks.\n " +
			"create - create a new task.\n" +
			"read - read task by id.\n" +
			"update - update an information in task.\n" +
			"delete - delete task by id.\n" +
			"done - mark as completed by id.\n ")
	}

	fmt.Println("Input id of the task. ")
	var inputstr string

	_, err = fmt.Scanf("%s", &inputstr)
	if err != nil {
		fmt.Println(err)
	}
	id, e := strconv.Atoi(inputstr)
	if e != nil {
		fmt.Println(e)
	}


	res, err := c.Read(context.Background(), &api.TaskId{Id: int32(id)})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Re %s",res.GetBody())
}