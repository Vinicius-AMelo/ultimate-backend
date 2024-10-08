FROM golang:1.23

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN go get -u github.com/air-verse/air
RUN go install github.com/air-verse/air
RUN mkdir -p tmp && chmod 777 tmp

RUN go build -v -o /usr/local/bin/app ./api/cmd

EXPOSE 8080

CMD ["air"]