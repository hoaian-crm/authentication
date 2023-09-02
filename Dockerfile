FROM golang:1.20.6

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest # install air to hot reload

COPY . ./

CMD [ "air" ]

EXPOSE 8080
