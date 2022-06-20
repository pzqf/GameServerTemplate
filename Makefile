.PHONY: all game win win_run
all:
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/game_server ./GameServer/main.go
	#mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/master_server ./MasterServer/main.go

game:
	mkdir -p bin/ &&env GOOS=linux GOARCH=amd64 go build  -o ./bin/game_server ./GameServer/main.go

win:
	go build  -o ./bin/game_server.exe ./GameServer/main.go

win_run:
	.\bin\game_server.exe -c .\config\GameServer.conf