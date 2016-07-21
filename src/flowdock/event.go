package flowdock
/*

message:
{"thread_id":"clHDVdAezeyVeGNEkbPd0FKqOFx",
 "event":"message",
 "tags":[],
 "uuid":"FHpQ4zXT2mHS3mCv",
 "thread":{"body":"",
           "external_url":null,
           "external_comments":0,
           "activities":0,
           "id":"clHDVdAezeyVeGNEkbPd0FKqOFx",
           "flow":"e494e46d-6b6a-4932-8c8a-6efdf0c12619",
           "status":null,
           "fields":[],
           "actions":[],
           "created_at":1469027127031,
           "title":"KAFFE-STATE",
           "updated_at":"2016-07-20T15:05:27.031Z",
           "initial_message":649,"internal_comments":1
          },
 "id":649,
 "flow":"e494e46d-6b6a-4932-8c8a-6efdf0c12619",
 "content":"KAFFE-STATE",
 "sent":1469027127031,
 "app":"chat",
 "created_at":"2016-07-20T15:05:27.031Z",
 "attachments":[],
 "user":"191106"
}

 */

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
