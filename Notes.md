# Notes

Install Go packages

```sh
$ go get -u github.com/gin-gonic/gin
go get -u github.com/line/line-bot-sdk-go/v7/linebot
go get -u github.com/spf13/viper
go get -u github.com/spf13/cobra@latest
go get -u go.mongodb.org/mongo-driver/mongo
```

## BSON
```
D: An ordered representation of a BSON document (slice)
M: An unordered representation of a BSON document (map)
A: An ordered representation of a BSON array
E: A single element inside a D type
```

* https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/

## MongoDB operator update
* `$setOnInsert` : https://www.mongodb.com/docs/v4.4/reference/operator/update/setOnInsert/