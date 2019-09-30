FROM golang:1.11.2

COPY . .
RUN go get -d github.com/gorilla/mux

CMD ["go","run","main.go"]

EXPOSE 9000