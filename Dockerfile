FROM golang:1.20.6

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest # install air to hot reload
RUN go install github.com/swaggo/swag/cmd/swag@latest # install swag command to open api

COPY . ./

CMD [ "air" ]

EXPOSE 8080
