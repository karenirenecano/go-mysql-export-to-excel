package export

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/karenirenecano/go-mysql-export-to-excel/task"
)

//Build the excel file
func Build(tasks task.Collection) {
	headers := map[string]string{
		"A1": "Task ID",
		"B1": "Title",
		"C1": "Is Completed",
		"D1": "Created At",
		"E1": "Updated At"}
	file := excelize.NewFile()
	for k, v := range headers {
		file.SetCellValue("Sheet1", k, v)
	}
	fmt.Println(tasks[1])
	for i := 0; i < len(tasks); i++ {
		appendRow(file, i, tasks)
	}

	var filename string = fmt.Sprintf("/tmp/tasks.xlsx")
	err := file.SaveAs(filename)
	if err != nil {
		fmt.Println(err)
	}
}

//Append every row, we can add styles if neeeded
func appendRow(file *excelize.File, index int, tasks task.Collection) (fileWriter *excelize.File) {
	rowCount := index + 2
	file.SetCellValue("Sheet1", fmt.Sprintf("A%v", rowCount), tasks[index].ID)
	file.SetCellValue("Sheet1", fmt.Sprintf("B%v", rowCount), tasks[index].Title)
	var isCompleted string = "In progress"
	if tasks[index].IsComplete == true {
		isCompleted = "Yes"
	}
	file.SetCellValue("Sheet1", fmt.Sprintf("C%v", rowCount), isCompleted)
	file.SetCellValue("Sheet1", fmt.Sprintf("D%v", rowCount), tasks[index].CreatedAt)
	file.SetCellValue("Sheet1", fmt.Sprintf("E%v", rowCount), tasks[index].UpdatedAt)

	return file

}
