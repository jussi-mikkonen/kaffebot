package flowdock
/*

{ event: 'file', content: { data: base64image, content_type: 'image/jpg', file_name: 'kahvi-status.jpg'}, thread_id: threadId}

 */

// Stupid hack ?
type Message interface {
	IsMessage()
}


type MessageFile struct {
	Event       string  `json:"event"`
	ThreadId    string  `json:"thread_id"`
	Content     struct{
		Data        string  `json:"data"`
		ContentType string  `json:"content_type"`
		FileName    string  `json:"file_name"`
	}                   `json:"content"`
}

func (mf *MessageFile) IsMessage() {
}


type MessageText struct {
	Event       string  `json:"event"`
	ThreadId    string  `json:"thread_id"`
	Content     string  `json:"content"`
}

func (mt *MessageText) IsMessage() {
}

