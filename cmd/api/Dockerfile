FROM golang:alpine as builder
ENV APPDIR $GOPATH/src/github.com/shirakiyo/ConveniGacha
ENV GO111MODULE on
RUN \
  apk update --no-cache && \
  mkdir -p $APPDIR
ADD . $APPDIR/
WORKDIR $APPDIR
RUN go mod tidy
RUN go build -ldflags "-s -w" -o conveni-api cmd/api/main.go
RUN mv conveni-api /

FROM alpine
RUN apk add --no-cache ca-certificates
RUN apk add mysql-client
COPY --from=builder /conveni-api ./
ENTRYPOINT ["./conveni-api"]

