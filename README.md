# linebot-go
Sample LineBot. receive message from line webhook, save the user info and message in MongoDB.

Use below golang libs:
* HTTP framework: https://github.com/gin-gonic/gin
* Config: https://github.com/spf13/viper
* Mongo driver: https://github.com/mongodb/mongo-go-driver 
* cobra: command line tools https://github.com/spf13/cobra
* Go line sdk : https://github.com/line/line-bot-sdk-go

## Ngrok
Download : https://dashboard.ngrok.com/get-started/setup

Connect your account
`ngrok config add-authtoken <your auth>`

Run

`./ngrok http <port>`

```sh
./ngrok http 8080
```

## Run locally
### Docker
```sh
docker-compose up --build

# 重新執行且放到背景
docker-compose start

# 停止容器
docker-compose stop
```

### Command
```sh
# command 
# go run main.go --config <your config> http --port <your port>
go run main.go --config config.toml http --port 8080

# air -c  <your air config> --port <your port>
# will run local.toml
air -c air.toml http

# build/
./build/linebot-go.mac.x64 --config local.toml http --port 8080
```

### Test Http Server
```sh
curl 127.0.0.1:8080/ping
```

## Line Webhook settings
```sh
# <ngrok url>/api/v1/callback
https://835d-61-228-16-110.jp.ngrok.io/api/v1/callback
```

## Connect Mongodb
* Root Username = root
* Root Password = root
* User Username = user
* User Password = user_password
* Database = linebot-go

# Reference
* [developers.line.biz](https://developers.line.biz/en/docs/)
* [使用 Docker 構築不同 MongoDB 架構 (二) - Standalone](https://ithelp.ithome.com.tw/articles/10224871)
* [使用Docker建立Mongodb加上Mongo Express](https://104.es/2022/07/05/docker-compose-mongodb-mongo-express/)
* [How to create a DB for MongoDB container on start up?](https://stackoverflow.com/questions/42912755/how-to-create-a-db-for-mongodb-container-on-start-up)
* [mongodb : db.getSiblingDB()](https://www.mongodb.com/docs/manual/reference/method/db.getSiblingDB/)
* [MONGO_INITDB_DATABASE and directory docker-entrypoint-initdb.d not running](https://github.com/docker-library/mongo/issues/429)
* [Docker Compose MongoDB docker-entrypoint-initdb.d is not working](https://stackoverflow.com/questions/60522471/docker-compose-mongodb-docker-entrypoint-initdb-d-is-not-working)
* [config: spf13/viper](https://github.com/spf13/viper)
* [cosmtrek/air](https://github.com/cosmtrek/air)