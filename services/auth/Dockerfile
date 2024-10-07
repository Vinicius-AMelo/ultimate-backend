FROM golang:1.23

WORKDIR /usr/src/app

COPY ./services/auth/go.mod ./services/auth/go.sum ./
RUN go mod download && go mod verify

COPY ./services/auth .

RUN go get -u github.com/air-verse/air
RUN go install github.com/air-verse/air
RUN mkdir -p tmp && chmod 777 tmp

RUN go build -v -o /usr/local/bin/app ./cmd

EXPOSE 8081


CMD ["air"]