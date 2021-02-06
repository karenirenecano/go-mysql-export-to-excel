package task

import (
	"github.com/karenirenecano/go-mysql-export-to-excel/database"
)

//GetAllByUserID Take in a userID and get all its tasks
//Return value will be a slice of Task, Error
func GetAllByUserID(userID int) ([]Task, error) {
	db := database.DBInstance()
	query := `SELECT * from tasks where user_id = ?;`
	results, err := db.Query(query, userID)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	tasks := make([]Task, 0)
	for results.Next() { // foreach row
		var task Task
		// The position of the columns should be the same that of select query above.
		// It should also follow the struct succession
		results.Scan(
			&task.ID,
			&task.UserID,
			&task.Title,
			&task.IsComplete,
			&task.CreatedAt,
			&task.UpdatedAt)
		tasks = append(tasks, task)
	}

	return tasks, nil
}
