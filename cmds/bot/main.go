package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/vanhtuan0409/gotokyo/pkg"
)

var (
	// config
	addr            string
	name            string
	eventBufferSize uint64
	debug           bool

	// global data
	botID int
)

func main() {
	flag.StringVar(&addr, "addr", "tokyo.thuc.space", "Tokyo server address")
	flag.StringVar(&name, "name", "atv", "Bot name")
	flag.Uint64Var(&eventBufferSize, "event-buffer", 10000, "Event buffer size")
	flag.BoolVar(&debug, "debug", false, "Debug log")
	flag.Parse()
	if name == "" {
		panic("Invalid bot name")
	}

	wsAddr := fmt.Sprintf("ws://%s/socket?key=%s&name=%s", addr, name, name)
	conn, _, err := websocket.DefaultDialer.Dial(wsAddr, nil)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	source := pkg.NewEventSource(conn, debug)
	stream := source.Stream(context.Background(), uint(eventBufferSize))

	initEvent := <-stream
	if err := initEvent.UnmarlshalData(&botID); err != nil {
		panic(err)
	}
	log.Printf("Bot id: %d", botID)

	var state pkg.GameState
	for {
		updateEvent := <-stream
		if err := updateEvent.UnmarlshalData(&state); err != nil {
			break
		}

		log.Printf("Event: %+v", state.Bounds)
	}
}
