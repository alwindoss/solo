package solo

import "fmt"

// Bot defines the functions a bot has to implement
type Bot interface {
	Run()
}

// NewBot creates a new bot
func NewBot(name string) Bot {
	switch name {
	case "echo":
		return &echoBot{}
	default:
		return &echoBot{}
	}
}

type echoBot struct {
}

func (echoBot) Run() {
	fmt.Printf("Echo....\n")
}
