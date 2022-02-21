FROM golang:1.11

WORKDIR /go/src/github.com/dmitryovchinnikov/rest_api
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# ADD . /resp_api

RUN  go build -o rest_api .

EXPOSE 8000:8000

CMD ["./rest_api"]