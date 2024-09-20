FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./services/auth/cmd

EXPOSE 8081

CMD ["/usr/local/bin/app"]