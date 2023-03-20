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

## Header
* https://docs.gitlab.com/ee/api/rest/#other-pagination-headers

```sh
x-next-page	The index of the next page.
x-page	The index of the current page (starting at 1).
x-per-page	The number of items per page.
x-prev-page	The index of the previous page.
x-total	The total number of items.
x-total-pages	The total number of pages.
```