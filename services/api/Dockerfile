FROM golang:1.23

WORKDIR /usr/src/app

COPY ./services/api/go.mod ./services/api/go.sum ./
RUN go mod download && go mod verify

COPY ./services/api .

RUN go get -u github.com/air-verse/air
RUN go install github.com/air-verse/air
RUN mkdir -p tmp && chmod 777 tmp

RUN go build -v -o /usr/local/bin/app ./cmd

EXPOSE 8080

CMD ["air"]