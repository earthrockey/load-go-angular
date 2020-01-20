FROM golang:1.13.5

WORKDIR /go/src/github.com/earthrockey/dc-go-ng-ma/backend
COPY . .
RUN go get github.com/rs/cors
RUN go get github.com/julienschmidt/httprouter
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]
