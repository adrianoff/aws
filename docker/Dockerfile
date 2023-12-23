FROM ubuntu:22.04

WORKDIR /usr/src/aws-server

RUN apt-get update && \
    apt-get -y install python3 python3-dev python3-pip golang-go apt-utils nano && \
    ln -s /usr/bin/python3 /usr/bin/python

RUN pip3 install --upgrade pip
COPY ./converter/requirements.txt /usr/src/aws-server/requirements.txt
RUN pip3 install -r /usr/src/aws-server/requirements.txt

COPY ./converter/convert.py /usr/src/aws-server
COPY ./server/server.go /usr/src/aws-server

RUN go build server.go
RUN rm server.go
RUN rm requirements.txt