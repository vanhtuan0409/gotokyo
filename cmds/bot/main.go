package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/vanhtuan0409/gotokyo/pkg"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	conf, err := ParseConfig()
	if err != nil {
		panic(err)
	}
	log.Printf("Conf: %+v", conf)

	client, err := pkg.NewClient(pkg.ClientOpt{
		Addr:       conf.Server.Address,
		BotName:    conf.Bot.Name,
		DebugRead:  conf.Client.DebugRead,
		DebugWrite: conf.Client.DebugWrite,
	})
	if err != nil {
		panic(err)
	}
	defer client.Close()

	source := pkg.NewEventSource(client, pkg.SourceOpt{
		Rate: conf.Client.EventRate,
	})
	behaviour := randomBehaviour{t: time.Now()}
	bot := pkg.NewBot(conf.Bot.Name, &behaviour, client)
	loop := pkg.NewGameLoop()
	loop.Run(context.Background(), bot, source)
}

type randomBehaviour struct {
	t time.Time
}

func (b *randomBehaviour) Process(tick uint64, bot *pkg.Bot, state *pkg.GameState) error {
	ctx := context.Background()
	bot.AdjustSpeed(ctx, rand.Float32())
	bot.RotateDeg(ctx, rand.Intn(360))
	bot.Fire(ctx)

	// if time.Since(b.t) > time.Duration(5*time.Second) {
	// 	bot.ChangeBehaviour(&standStillBebaviour{})
	// }

	return nil
}

type standStillBebaviour struct{}

func (b *standStillBebaviour) Process(tick uint64, bot *pkg.Bot, state *pkg.GameState) error {
	bot.AdjustSpeed(context.Background(), 0)
	return nil
}
