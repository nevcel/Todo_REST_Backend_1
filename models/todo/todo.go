// Package todo contains the model parts
package todo

import "strconv"

// Todo type definition with json tags
type Todo struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Terminated  bool   `json:"terminated"`
}

// Serialize serializes the passed to todo into a slice form
func (t Todo) Serialize() []string {
	todoSerialized := []string{t.Id, t.Title, t.Description, strconv.FormatBool(t.Terminated)}
	return todoSerialized
}
