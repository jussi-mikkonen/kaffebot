package flowdock


const (
	EventMessage        string = "message"
	EventFile           string = "file"
	EventUserActivity   string = "activity.user"
)

type Event struct {
	// TODO Add more!
	ThreadId        string      `json:"thread_id"`
	Event           string      `json:"event"`
	Content         string      `json:"content"`
}
