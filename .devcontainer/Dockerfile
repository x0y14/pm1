FROM ubuntu:22.04

# install go
RUN apt update -y -q && apt upgrade -y -q
RUN DEBIAN_FRONTEND=noninteractive apt install --no-install-recommends -y -q curl build-essential ca-certificates git 
RUN curl https://go.dev/dl/go1.19.2.linux-amd64.tar.gz | tar xvf -C /usr/local
ENV PATH $PATH:/usr/local/go/bin

# install gtk4
RUN apt install libgtk-4-dev

WORKDIR /workspace
ADD . /workspace/

