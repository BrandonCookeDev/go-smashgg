FROM golang:1.8

WORKDIR /go/src/gosmashgg
COPY . .

RUN go get -u github.com/tidwall/gjson

RUN go install
WORKDIR /go/src/gosmashgg/samples
CMD ["go", "run", "Index.go"]