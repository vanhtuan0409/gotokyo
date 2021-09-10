package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/vanhtuan0409/gotokyo/pkg"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var (
	// config
	addr            string
	name            string
	eventBufferSize uint64
	debug           bool
)

func main() {
	flag.StringVar(&addr, "addr", "ws://localhost:8091", "Tokyo server address")
	flag.StringVar(&name, "name", "atv", "Bot name")
	flag.Uint64Var(&eventBufferSize, "event-buffer", 10000, "Event buffer size")
	flag.BoolVar(&debug, "debug", false, "Debug log")
	flag.Parse()

	client, err := pkg.NewClient(pkg.ClientOpt{
		Addr:    addr,
		BotName: name,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	source := pkg.NewEventSource(client, pkg.SourceOpt{})

	bot := pkg.NewBot(name, &randomBehaviour{
		t: time.Now(),
	}, client)
	bot.Run(
		context.Background(),
		source,
	)
}

type randomBehaviour struct {
	t time.Time
}

func (b *randomBehaviour) Process(bot *pkg.Bot, state *pkg.GameState) error {
	log.Printf("%+v", state.Bounds)

	ctx := context.Background()
	bot.AdjustSpeed(ctx, rand.Float32())
	bot.RotateDeg(ctx, rand.Intn(180))
	bot.Fire(ctx)

	if time.Since(b.t) > time.Duration(5*time.Second) {
		bot.ChangeBehaviour(&standStillBebaviour{})
	}
	return nil
}

type standStillBebaviour struct{}

func (b *standStillBebaviour) Process(bot *pkg.Bot, state *pkg.GameState) error {
	bot.AdjustSpeed(context.Background(), 0)
	return nil
}
