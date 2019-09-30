FROM golang:alpine

COPY . /go/src/github.com/PetIsland/config-service

WORKDIR /go/src/github.com/PetIsland/config-service

RUN go get .
RUN go build

CMD ["/go/bin/config-service"]

EXPOSE 8888
