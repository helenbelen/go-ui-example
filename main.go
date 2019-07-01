package main

import (
	"log"

	"github.com/Equanox/gotron"
)

type CustomEvent struct {
	*gotron.Event
	CustomAttribute string `json:"AtrNameInFrontend"`
}

func main() {
	window, err := gotron.New("www")
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 720
	window.WindowOptions.Height = 720
	window.WindowOptions.Title = "GoTron"

	onEvent := gotron.Event{Event: "click"}

	window.On(&onEvent, func(bin []byte) {
		log.Println("received hello")
		log.Println(bin)

	})

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	<-done
}
