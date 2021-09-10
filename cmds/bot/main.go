package main

import (
	"context"
	"flag"
	"log"
	"math"

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

	client, err := pkg.NewClient(pkg.ClientOpt{
		Addr:    addr,
		BotName: name,
		Debug:   true,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	source := pkg.NewEventSource(client, pkg.SourceOpt{
		Rate: 1000,
	})

	go debugEvent(source)
	debugCommand(client)
}

func debugEvent(source *pkg.EventSource) {
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

func debugCommand(client *pkg.Client) {
	for {
		if err := client.SendCommand(pkg.RotateCommand{
			Angle: math.Pi / 6,
		}); err != nil {
			log.Printf("Command err: %+v", err)
		}
		if err := client.SendCommand(pkg.FireCommand{}); err != nil {
			log.Printf("Command err: %+v", err)
		}
	}
}
