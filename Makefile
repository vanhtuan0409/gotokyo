run:
	go run ./cmds/bot

run_live:
	go run ./cmds/bot -server.addr "ws://tokyo.thuc.space" -bot.name "tuan"

server:
	docker run -it -p 8091:8080 -e RUST_BACKTRACE=1 ledongthuc/tokyo-rs:latest
