FROM golang:1.12 as builder

LABEL maintainer="Pavel Popov pashkaposhka42@gmail.com"

RUN mkdir /app

ADD . /app/

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -o /main .

FROM alpine:3.9

RUN apk --no-cache add ca-certificates && \
    apk update && \
    mkdir /app

WORKDIR /app

COPY --from=builder /main ./

CMD ["./main"]
