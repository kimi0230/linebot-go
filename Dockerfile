# build stage
FROM golang:1.20-alpine AS build-env
WORKDIR '/go/src/linebot-go/api'
ADD . .
RUN apk add git
RUN go build -o APPSERVICE

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/linebot-go/api/APPSERVICE /app/
COPY --from=build-env /go/src/linebot-go/api/config.toml /app/
# COPY --from=build-env /go/src/linebot-go/api/.env /app/
# RUN ["mkdir","logs"]
RUN ["chmod", "+x", "APPSERVICE"]
RUN ["ls", "-al"]
EXPOSE 8080

# ENTRYPOINT ./APPSERVICE
CMD ./APPSERVICE http 8080