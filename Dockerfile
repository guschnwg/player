FROM golang:1.15.7 AS base

ENV PORT 8000

RUN apt-get update
RUN apt-get install -y gzip

COPY . /app
WORKDIR /app

RUN make client
RUN make client_compress_gzip
RUN make server


FROM python:3.8-buster AS dist

COPY --from=base /app/app.out /
COPY --from=base /app/web /web

RUN pip install youtube_dl

CMD PORT=${PORT} USE_GZIP=true ./app.out