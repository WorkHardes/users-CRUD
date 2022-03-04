FROM golang:1.16-alpine

COPY . /users-CRUD

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/users/main.go

WORKDIR /users-CRUD

CMD [ "./app" ]
