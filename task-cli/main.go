package main

func main() {
	todos := Todos{}
	Storage := NewStorage[Todos]("first-todos.json")
	Storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)

	todos.Print()
	Storage.Save(todos)
}
