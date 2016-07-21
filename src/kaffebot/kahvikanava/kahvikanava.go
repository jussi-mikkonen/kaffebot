package kahvikanava

import (
	"flowdock"
	"strings"
	"sync"
	"log"
	"os"
	"encoding/base64"
	"io/ioutil"
	"net"
)

const (
	KAFFEFILE       = "/home/pi/image.jpg"
	//KAFFEFILE = "/home/jussi/kahvitesti.jpg"
	KAFFE_COMMAND   = "KAFFE-STATE"
	IP_COMMAND      = "IP-STATE"
)
var logger *log.Logger = log.New(os.Stdout, "KAHVIKANAVA -> ", 0)

// Initialisation
func KaffeKanavaBot(wg *sync.WaitGroup) {

	// We want kaffebot to reconnect on disconnect / error / etc
	for {
		var flow *flowdock.Flowdock = flowdock.New()


		flow.SetLogger(logger)

		flow.ApiKey = os.Getenv("KAFFEBOT_KAFFE_FLOW_API_KEY")
		flow.FlowName = os.Getenv("KAFFEBOT_KAFFE_FLOW_FLOWNAME")
		flow.CompanyName = os.Getenv("KAFFEBOT_KAFFE_FLOW_COMPANY")

		// Need another functionality to KahviKanava?
		// Add a new handler here.
		logger.Println("Registering KaffeHandler...")
		var handler *KaffeHandler = &KaffeHandler{}
		flow.RegisterHandler(handler)

		logger.Println("Registering IPHandler")
		var ipHandler *IPHandler = &IPHandler{}
		flow.RegisterHandler(ipHandler)

		logger.Println("Starting flowdock.Connect() for KahviKanava")
		flow.Connect()
	}

	wg.Done()
}

type IPHandler struct {

}

func (handler IPHandler) OnEvent(event flowdock.Event) flowdock.Message  {

	if(event.Event == flowdock.EventMessage) && strings.Contains(event.Content, IP_COMMAND) {
		message := &flowdock.MessageText{Event: flowdock.EventMessage }
		message.ThreadId = event.ThreadId
		message.Content = "Virhe!"

		addrs, err := net.InterfaceAddrs()

		if err != nil {
			logger.Println("Failed call to net.InterfaceAddrs()")
			return message
		}
		message.Content = "IP osoitteet: "
		for _, addr := range addrs {
			message.Content = message.Content + addr.String() + "   "
		}

		return message
	}

	return nil
}



// Handler type for KAFFE-STATE command
type KaffeHandler struct {
}

// EventHandler for events from flow. Implements interface flowdock.EventHandler.
func (handler KaffeHandler) OnEvent(event flowdock.Event) flowdock.Message {

	if (event.Event == flowdock.EventMessage) && (strings.Contains(event.Content, KAFFE_COMMAND)) {

		message := &flowdock.MessageFile{Event: flowdock.EventFile}
		message.ThreadId = event.ThreadId
		message.Content.ContentType = "image/jpeg"
		message.Content.FileName = "kahvi-status.jpg"
		message.Content.Data = readFileToBase64String(KAFFEFILE)

		if message.Content.Data == "" {
			return nil
		}

		return message
	}

	return nil
}

func readFileToBase64String(fileName string) string {

	encoding := base64.StdEncoding
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return ""
	}

	return encoding.EncodeToString(data)
}
