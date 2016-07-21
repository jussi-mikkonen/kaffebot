package flowdock

type EventHandler interface {
	OnEvent(Event) Message
}
