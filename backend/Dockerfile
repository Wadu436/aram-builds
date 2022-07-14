FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY cmd ./cmd
COPY db ./db
COPY server ./server
COPY main.go ./

RUN go build -o ./backend

FROM debian:buster-slim as runner

WORKDIR /app

COPY --from=build /app/backend backend

CMD [ "./backend", "server" ]