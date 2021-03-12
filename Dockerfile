FROM golang:latest

LABEL maintainer="amartery@gmail.com"

WORKDIR $GOPATH/src/app
COPY . $GOPATH/src/app

RUN go get -d -v ./...
RUN go install -v ./...

RUN make build

EXPOSE 8080

CMD ["./main"]