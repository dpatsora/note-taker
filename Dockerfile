FROM golang:1.21.3-alpine AS builder

WORKDIR /build
ENV CGO_ENABLED=1

RUN apk add build-base curl

COPY . .

RUN go build -o ./bin/note-taker .

FROM alpine

RUN addgroup -g 1001 -S note-taker
RUN adduser -S note-taker -u 1001 -G note-taker

USER note-taker

COPY --from=builder /build/bin /bin

CMD ["note-taker", "start"]