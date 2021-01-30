FROM golang:alpine

COPY go.mod /usr/src/app/
COPY go.sum /usr/src/app/

WORKDIR /usr/src/app/api/
RUN go mod download

COPY main.go /usr/src/app/api/
COPY mongo.go /usr/src/app/api/
COPY person.go /usr/src/app/api/

RUN go build -o /usr/src/app/bin/api .

WORKDIR /usr/src/app/bin/

CMD ./api