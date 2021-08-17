package entities

type Task struct {
	ID        uint64 `json:"id,omitempty"`
	Name      string `json:"name"`
	Timestamp int64  `json:"timestamp"`
}
