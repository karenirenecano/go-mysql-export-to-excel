package main

import (
	"fmt" //standard libraries, out of the box

	"github.com/karenirenecano/go-mysql-export-to-excel/task"
	"github.com/karenirenecano/go-mysql-export-to-excel/utility" //your custom local package, go.mod file
)

//This is the main entry point of your go code
func main() {
	utility.Setup()
	var userID int = 2
	tasks, err := task.GetAllByUserID(userID)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(tasks)
}
