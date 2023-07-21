FROM golang:alpine as builder

RUN apk add git
RUN apk add make

WORKDIR /app

RUN git clone -b main --depth 1 https://github.com/lizongying/go-crawler.git
RUN cd go-crawler && make testSpider

FROM alpine:latest as prod

WORKDIR /app

COPY --from=builder /app/go-crawler/releases/testSpider /usr/bin/spider
COPY --from=builder /app/go-crawler/example.yml .

#ENTRYPOINT ["spider"]