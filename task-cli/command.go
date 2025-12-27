package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Add    string
	Del    int
	Edit   string
	Status string
	List   bool
}

func NewCmdFlags() *Command {
	cf := Command{}

	flag.StringVar(&cf.Add, "add", "", "add todo")
	flag.StringVar(&cf.Edit, "update", "", "update task name")
	flag.StringVar(&cf.Status, "status", "", "change task status")
	flag.IntVar(&cf.Del, "delete", -1, "delete task by id")
	flag.BoolVar(&cf.List, "list", false, "print all task")

	flag.Parse()

	return &cf
}

func (cf *Command) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.List:
		todos.Print()
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format. Please use id:new_description")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Invalid Index")
			os.Exit(1)
		}

		todos.update(parts[1], index)
	case cf.Status != "":
		parts := strings.SplitN(cf.Status, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error, invalid format. Please use id:new_status")
			os.Exit(1)
		}

		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Invalid Index")
			os.Exit(1)
		}

		todos.StatusChange(parts[1], index)
	case cf.Del != -1:
		todos.delete(cf.Del)
	default:
		fmt.Println("Invalid Command")
	}
}
