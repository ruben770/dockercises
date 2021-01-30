FROM golang:alpine

COPY go.mod /usr/src/app/
COPY go.sum /usr/src/app/

WORKDIR /usr/src/app/
RUN go mod download

COPY main.go /usr/src/app/populate/
COPY mongo.go /usr/src/app/populate/
COPY people.xml /usr/src/app/data/

WORKDIR /usr/src/app/populate/
RUN go build -o /usr/src/app/bin/populate .

WORKDIR /usr/src/app/bin

CMD ./populate
