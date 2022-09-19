FROM golang:alpine

RUN apk --no-cache add gcc g++ make git

WORKDIR /go/src/github.com/artarts36/potaynik

COPY . .

RUN make deps

RUN GOOS=linux go build -ldflags="-s -w" -o /go/bin/web-app ./cmd/main.go

##################################################################################################
#                                   Setup entrypoint                                             #
##################################################################################################

RUN apk --no-cache add ca-certificates

EXPOSE 80

ENTRYPOINT /go/bin/web-app
