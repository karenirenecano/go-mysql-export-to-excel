package task

//Task Struct
//Struct is a collection of data fields with declared data types
type Task struct {
	ID         int    `json:"id"`          //serves as the primary key of the task table
	UserID     int    `json:"user_id"`     //reference foreign key from the user table, owner of a specific task
	Title      string `json:"title"`       //title about the task
	IsComplete bool   `json:"is_complete"` //is the task completed already
	CreatedAt  string `json:"created_at"`  //created timestamp
	UpdatedAt  string `json:"updated_at"`  //updated timestamp
}

//Collection of task slice
type Collection []Task
