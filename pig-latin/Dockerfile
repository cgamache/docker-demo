FROM golang:latest

WORKDIR "/pig"
RUN go get github.com/stretchr/piglatin
COPY pig.go /pig
RUN go build

ENTRYPOINT ["/pig/pig"]