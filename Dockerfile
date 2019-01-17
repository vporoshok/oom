FROM golang

WORKDIR /go/src/oom
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT [ "oom" ]
