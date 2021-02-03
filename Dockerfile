FROM golang:1.15.7 AS base

ENV PORT 8000

RUN apt-get update
RUN apt-get install -y python3 python3-pip gzip
RUN pip3 install youtube_dl

COPY . /app
WORKDIR /app

RUN make client
RUN make client_compress_gzip
RUN make server

CMD PORT=${PORT} USE_GZIP=true ./app.out