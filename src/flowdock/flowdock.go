package flowdock

import (
	"log"
	"net/http"
	"encoding/json"
	"bufio"
	"bytes"
	"io/ioutil"
)

type Flowdock struct {
	ApiKey      string
	FlowName    string
	CompanyName string
	channel     chan Event
	handlers    []EventHandler
	logger      *log.Logger
}

func New() *Flowdock {
	c := make(chan Event)
	handlers := []EventHandler{}
	return &Flowdock{channel: c, handlers: handlers}
}

func (flow *Flowdock) SetLogger(paramLogg *log.Logger)  {
	flow.logger = paramLogg
}

func (flow *Flowdock) RegisterHandler(handler EventHandler) {
	flow.handlers = append(flow.handlers, handler)
}

func (flow Flowdock) Connect() (err error) {
	err = nil

	// Start reading events from flow in go-routine
	go flow.readerRoutine()

	// Read events sent from the readerRoutine
	for event := range flow.channel {
		for _, handler := range flow.handlers {
			if handler != nil { // should always be true (:
				message := handler.OnEvent(event)
				if message != nil {
					flow.Send(message)
				}
			}
		}
	}

	return
}

func (flow *Flowdock) readerRoutine() {
	client := http.Client{}

	req, reqErr := http.NewRequest("GET", "https://stream.flowdock.com/flows/" + flow.CompanyName + "/" + flow.FlowName, nil)
	if reqErr != nil {
		flow.logger.Println("Failed to create request.")
		return
	}
	req.SetBasicAuth(flow.ApiKey, "")

	response, respError := client.Do(req)
	if respError != nil {
		flow.logger.Println("Failed to do request!")
		return
	}

	reader := bufio.NewReader(response.Body)

	for {
		data, readErr := reader.ReadBytes('\n')

		if readErr != nil {
			flow.logger.Println("Read error from HTTP stream, breaking out of loop, stopping this bot.")
			break
		}

		var event Event
		json.Unmarshal(data, &event)
		flow.channel <- event
	}

	close(flow.channel)
}

func (flow *Flowdock) Send(message Message) error {
	json, err := json.Marshal(message)

	flow.logger.Println("Sending message to " + flow.FlowName)

	if err != nil {
		flow.logger.Println("Failed to marshal Message to json!")
		return nil //error.Error("Failed to marshal Message to json!")
	}
	client := http.Client{}

	postReq, err := http.NewRequest("POST", "https://api.flowdock.com/flows/" + flow.CompanyName + "/" + flow.FlowName + "/messages", bytes.NewReader(json))
	postReq.Header["Content-Type"] = append(postReq.Header["Content-Type"], "application/json")

	if err != nil {
		flow.logger.Println("Failed to create POST request")
		return nil // TODO error
	}
	postReq.SetBasicAuth(flow.ApiKey, "")

	resp, err := client.Do(postReq)

	if err != nil {
		flow.logger.Println("Failed to send the POST request!")
		return nil // TODO error
	}

	if resp.StatusCode != http.StatusCreated {
		flow.logger.Printf("GOT HTTP STATUS: %d\n", resp.StatusCode)
		flow.logger.Printf("Sent: %s", json)
		bodyData, _ := ioutil.ReadAll(resp.Body)
		bodyStr := string(bodyData)
		flow.logger.Println(bodyStr)

	} else {
		flow.logger.Println("Message sent to flow: " + flow.FlowName)
		ioutil.ReadAll(resp.Body)
	}

	resp.Body.Close()

	return nil
}

