FROM golang:1.12.5-alpine

RUN mkdir -p /go/src/Kickoff

WORKDIR /go/src/Kickoff

COPY . /go/src/Kickoff

RUN go install Kickoff

CMD ["/go/bin/Kickoff"]

EXPOSE 8080