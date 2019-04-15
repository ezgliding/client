FROM golang:1.12-stretch AS builder
WORKDIR /build
COPY . /build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ezgliding .

FROM alpine:latest
WORKDIR /
RUN apk --no-cache add ca-certificates
COPY --from=builder /build/ezgliding .
CMD ["/ezgliding"]
