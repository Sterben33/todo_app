package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strconv"
	"todo_app/pkg/api"
)

var input string

func GetId() (int32, error){



	//Getting id
	fmt.Println("Input id of the task. ")
	_, err1 := fmt.Scanf("%s", &input)
	if err1 != nil { return -1, err1 }
	id, err2 := strconv.Atoi(input)
	if err2 != nil { return -1, err2 }

	return int32(id), nil
}

func GetData() (string, string, error){

	reader := bufio.NewReader(os.Stdin)
	//Getting title
	fmt.Print("Input new title: ")
	title, _ := reader.ReadString('\n')

	//Getting body
	fmt.Print("Input new body: ")
	body, _ := reader.ReadString('\n')

	return title, body, nil
}
func main() {

	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := api.NewToDoClient(conn)
	b := true
	for b {
		fmt.Println("Choose an option:\n" +
			"list - get list of the tasks.\n" +
			"create - create a new task.\n" +
			"read - read task by id.\n" +
			"update - update an information in task.\n" +
			"delete - delete task by id.\n" +
			"done - mark as completed by id.\n" +
			"q - quit. ")

		_, err = fmt.Scanf("%s", &input)
		if err != nil {
			fmt.Println(err)
		}
		option := input
		fmt.Println()
		switch option{
			case "list":
				res, e := c.GetList(context.Background(), &api.Empty{})
				if e != nil {
					fmt.Println(err)
					continue
				}
				for task := range res.Task {
					fmt.Printf("%v %v %v %v\n",res.Task[task].GetId(),res.Task[task].GetTitle(),res.Task[task].GetBody(),res.Task[task].GetDone())
					fmt.Println("--------------------")
				}

			case "create":
				title, body, err := GetData()
				if err != nil {
					fmt.Println(err)
				}
				res, e := c.Create(context.Background(), &api.Data{Title: title, Body: body})
				if e != nil {
					fmt.Println(e)
					continue
				}
				fmt.Println(res.GetBody())
			case "read":
				id, err := GetId()
				if err != nil {
					fmt.Println(err)
					continue
				}
				res, e := c.Read(context.Background(), &api.TaskId{Id: id})
				if e != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(res.GetBody())
			case "update":
				id, er := GetId()
				if er != nil {
					fmt.Println(er)
					continue
				}
				title, body, err := GetData()
				if err != nil {
					fmt.Println(err)
					continue
				}
				res, e := c.Update(context.Background(), &api.Task{Id: id,Title: title, Body: body})
				if e != nil {
					fmt.Println(e)
					continue
				}
				fmt.Println(res)
			case "delete":
				id, err := GetId()
				if err != nil {
					fmt.Println(err)
					continue
				}
				_, e := c.Delete(context.Background(), &api.TaskId{Id: id})
				if e != nil {
					fmt.Println(e)
					continue
				}
				fmt.Println("Deleted successfully. ")
			case "done":
				id, err := GetId()
				if err != nil {
					fmt.Println(err)
					continue
				}
				res, e := c.MarkAsDone(context.Background(), &api.TaskId{Id: id})
				if e != nil {
					fmt.Println(e)
					continue
				}
				fmt.Printf("%v", res.GetBody())

			case "q":
				b = false
		}
		fmt.Println()
		fmt.Println()
		if !b {
			break
		}
	}

}