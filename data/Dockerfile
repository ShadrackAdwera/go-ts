# Build image

FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app/

WORKDIR /app

RUN CGO_ENABLED=0 go build -o dataService ./cmd/api

RUN chmod +x /app/dataService

# build a tiny docker image - copy over my executable file

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/dataService /app/

CMD [ "/app/dataService" ]