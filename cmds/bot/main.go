package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/vanhtuan0409/gotokyo/pkg"
	"github.com/vanhtuan0409/gotokyo/pkg/behaviours"
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
		BufferSize: conf.Client.EventBufferSize,
		Rate:       conf.Client.EventRate,
	})

	compositeBeh := behaviours.NewBehaviourComposite(
		behaviours.CombineMovingStrat(
			behaviours.StratMoveAntiGravity,
			behaviours.StratMoveDodgePerpendicular,
		),
		behaviours.StratShootNone,
		// behaviours.NewStratShootingThrottle(
		// 	15,
		// 	behaviours.StratShootNearest,
		// ),
	)

	bot := pkg.NewBot(conf.Bot.Name, compositeBeh, client)
	loop := pkg.NewGameLoop()
	loop.Run(context.Background(), bot, source)
}
