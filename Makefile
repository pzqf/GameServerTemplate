.PHONY: all game master
all:
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/game_server ./GameServer/main.go
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/master_server ./MasterServer/main.go

game:
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/game_server ./GameServer/main.go

master:
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/master_server ./MasterServer/main.go

