go build -o bin/serviceloader steve/serviceloader 

go build -o bin/room/room.so -buildmode=plugin steve/room
cp room/config.yml configs/room/config.yml 

go build -o bin/gateway/gateway.so -buildmode=plugin steve/gateway
cp gateway/config.yml configs/gateway/config.yml 


go build -o bin/match/match.so -buildmode=plugin steve/match 
cp match/config.yml configs/match/config.yml 

go build -o bin/login/login.so -buildmode=plugin steve/login 
cp login/config.yml configs/login/config.yml 

