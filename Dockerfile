FROM golang:1.17-alpine as builder
ENV GOOS=linux \
    GOARCH=arm64 \
    CGO_ENABLED=0

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod download && go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10
COPY --from=builder /go/bin/app /
ENTRYPOINT ["/app"]