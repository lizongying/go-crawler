FROM golang:alpine as builder

RUN apk add git
RUN apk add make

WORKDIR /app

RUN git clone -b main --depth 1 https://github.com/lizongying/go-crawler.git
RUN cd go-crawler && make test_spider

FROM alpine:latest as prod

WORKDIR /app

COPY --from=builder /app/go-crawler/releases/test_spider /usr/bin/spider
COPY --from=builder /app/go-crawler/example.yml .

#ENTRYPOINT ["spider"]