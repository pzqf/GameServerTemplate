@echo off
docker run -d \
	--restart=always \
	-e TZ=Asia/Shanghai \
	--name game_server \
	-p 9160:9160 \
	-v `pwd`/bin/:/work \
	-w /work \
	ubuntu \
	./game_server