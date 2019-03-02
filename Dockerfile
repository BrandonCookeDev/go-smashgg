FROM golang:1.8

WORKDIR /go/src/smashggo
COPY . .

RUN go get -u github.com/tidwall/gjson

RUN go build
CMD ["./smashggo"]