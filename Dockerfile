FROM golang:1.14.1 AS base

RUN apt-get update
RUN apt-get install -y python3 python3-pip
RUN pip3 install youtube_dl

COPY . /app
WORKDIR /app

RUN make clean
WORKDIR /app

RUN make client
WORKDIR /app

RUN make server
WORKDIR /app

CMD make run