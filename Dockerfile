#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./...

#final stage
FROM alpine:latest
EXPOSE 3000
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /zoneinfo.zip
ENV ZONEINFO /zoneinfo.zip
ENTRYPOINT /app
LABEL Name=housingbot Version=0.0.1
