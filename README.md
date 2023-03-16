# linebot-go

use below golang libs:
* HTTP framework: https://github.com/gin-gonic/gin
* Config: https://github.com/spf13/viper
* Mongo driver: https://github.com/mongodb/mongo-go-driver 
* cobra: command line tools https://github.com/spf13/cobra
* Go line sdk : https://github.com/line/line-bot-sdk-go


```sh
$ go get -u github.com/gin-gonic/gin
go get -u github.com/line/line-bot-sdk-go/v7/linebot
go get -u github.com/spf13/viper
go get -u github.com/spf13/cobra@latest
go get -u go.mongodb.org/mongo-driver/mongo
```

## Run locally
```sh
docker-compose up --build

# 重新執行且放到背景
docker-compose start

# 停止容器
docker-compose stop
```

# Reference
* [developers.line.biz](https://developers.line.biz/en/docs/)