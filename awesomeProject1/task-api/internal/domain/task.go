package domain

type Status string

const (
	StatusPending  Status = "pending"
	StatusComplete Status = "complete"
)

type Task struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}
