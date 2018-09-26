FROM golang:1.9.6
WORKDIR /go/src/github.com/mnuma/sample-go-worker/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/mnuma/sample-go-worker/main .
CMD ["./main"]
