package models

type Task struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	DueDate  string `json:"due_date"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
}
