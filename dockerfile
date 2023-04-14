FROM golang:1.20-alpine

WORKDIR /usr/src/userapi

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ENV HTTP_PORT=8080
ENV JSON_STORE_NAME=./store/users.json

RUN go build -o /usr/src/userapi/build/userapi /usr/src/userapi/cmd/http/main.go

CMD [ "/usr/src/userapi/build/userapi" ]
