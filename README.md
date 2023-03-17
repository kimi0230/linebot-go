# linebot-go

use below golang libs:
* HTTP framework: https://github.com/gin-gonic/gin
* Config: https://github.com/spf13/viper
* Mongo driver: https://github.com/mongodb/mongo-go-driver 
* cobra: command line tools https://github.com/spf13/cobra
* Go line sdk : https://github.com/line/line-bot-sdk-go

## Ngrok
Download
* https://dashboard.ngrok.com/get-started/setup


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
```
go run main.go -c local.toml
```

# Reference
* [developers.line.biz](https://developers.line.biz/en/docs/)
* [使用 Docker 構築不同 MongoDB 架構 (二) - Standalone](https://ithelp.ithome.com.tw/articles/10224871)
* [使用Docker建立Mongodb加上Mongo Express](https://104.es/2022/07/05/docker-compose-mongodb-mongo-express/)
* [How to create a DB for MongoDB container on start up?](https://stackoverflow.com/questions/42912755/how-to-create-a-db-for-mongodb-container-on-start-up)
* [mongodb : db.getSiblingDB()](https://www.mongodb.com/docs/manual/reference/method/db.getSiblingDB/)
* [MONGO_INITDB_DATABASE and directory docker-entrypoint-initdb.d not running](https://github.com/docker-library/mongo/issues/429)
* [Docker Compose MongoDB docker-entrypoint-initdb.d is not working](https://stackoverflow.com/questions/60522471/docker-compose-mongodb-docker-entrypoint-initdb-d-is-not-working)
* [config: spf13/viper](https://github.com/spf13/viper)