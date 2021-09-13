run:
	go run ./cmds/bot

run_live:
	go run ./cmds/bot -server.addr "ws://tokyo.thuc.space" -bot.name "coward_dog"

server:
	docker run -it -p 8091:8080 -e RUST_BACKTRACE=1 ledongthuc/tokyo-rs:latest

build:
	CGO_ENABLED=0 go build -o bins/${NAME}_bot ./cmds/bot

build_pi:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o bins/${NAME}_bot_arm ./cmds/bot
