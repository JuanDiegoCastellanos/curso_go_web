package main

import (
	"fmt"
)

// "bufio"

// "os"
// "strconv"
// "strings"
type task struct {
	name        string
	description string
	completed   bool
}

func (t *task) markComplete() {
	t.completed = true
}
func (t *task) updateDescription(description string) {
	t.description = description
}
func (t *task) updateName(newName string) {
	t.name = newName
}

type taskList struct {
	tasks []*task
}

func (tl *taskList) addElement(newtask *task) {
	tl.tasks = append(tl.tasks, newtask)
}

func (tl *taskList) deleteElement(index int) {
	// [:index] Desde el inicio hasta el índice
	// [index+1:] Desde el índice hasta el final
	tl.tasks = append(tl.tasks[:index], tl.tasks[index+1:]...)
}

// for en struct
func (tl *taskList) printTasks() {
	for _, task := range tl.tasks {
		fmt.Println("Name: ", task.name)
		fmt.Println("Description: ", task.description)
	}
}

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()

	// operation := scanner.Text()
	// fmt.Println(operation)
	// values := strings.Split(operation, "+")
	// fmt.Println(values)
	// fmt.Println(values[0] + values[1])

	// value1, _ := strconv.Atoi(values[0])
	// value2, _ := strconv.Atoi(values[1])

	// fmt.Println(value1 + value2)
	t := &task{
		name:        "Cut the wires",
		description: "Cortar los cables para que funcione normal el internet",
		completed:   false,
	}
	fmt.Println(t)
	fmt.Printf("%+v \n", t)
	t.markComplete()
	fmt.Printf("Con punteros --> %+v \n", t)

	t.updateName("Sing for the moment")
	fmt.Println("Sin punteros --> ", t.name)

	t1 := &task{
		name:        "Tarea 2",
		description: "Description tarea 2",
		completed:   false,
	}
	t2 := &task{
		name:        "Tarea 3",
		description: "Description tarea 3",
		completed:   false,
	}

	tasks := &taskList{
		tasks: []*task{t, t1},
	}

	fmt.Printf("%+v ", tasks.tasks[0])

	tasks.addElement(t2)

	fmt.Println(len(tasks.tasks))

	tasks.deleteElement(1)

	fmt.Println(len(tasks.tasks))

	// Iterar con for
	for i := 0; i < len(tasks.tasks); i++ {
		fmt.Println(tasks.tasks[i])
	}

	// range
	for index, task := range tasks.tasks {
		fmt.Println("index: ", index, " Task: ", task.name)
	}
	tasks.printTasks()

	// maps

	myMap := make(map[int]string)

	myMap[2022] = "Juan Diego Castellanos"

	fmt.Println(myMap[2022])

	tasksMap := make(map[string]*taskList)

	tasksMap["Juan Diego"] = tasks

	fmt.Println(tasksMap["Juan Diego"])

	var x int = 2

	println(&x)

}
