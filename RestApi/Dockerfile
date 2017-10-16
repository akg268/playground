FROM golang:1.8

ENV http_proxy http://proxy.kdc.capitalone.com:8099
ENV https_proxy http://proxy.kdc.capitalone.com:8099 

WORKDIR /go/src/app

COPY ./src/github.com/arunkumargan/main/main.go .
COPY ./src/github.com/arunkumargan/main/eil-qaid.key .
COPY ./src/github.com/arunkumargan/main/eil-qaid.pem .
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."
EXPOSE 8443

CMD ["go-wrapper", "run"] # ["app"]
