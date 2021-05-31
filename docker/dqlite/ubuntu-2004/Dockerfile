FROM ubuntu:20.04 as builder

ARG DEBIAN_FRONTEND=noninteractive

RUN apt-get clean  && \
    apt-get update && \
    apt-get -y install init \
            software-properties-common && \
    apt-get -y autoremove

RUN apt-get -y install \
    gcc \
    vim \
    curl \
    make \
    git

RUN curl -LSs https://dl.google.com/go/go1.16.linux-amd64.tar.gz -o go.tar.gz && \
    tar -xf go.tar.gz && \
    rm -v go.tar.gz && \
    mv go /usr/local/

WORKDIR /utils
RUN git clone --depth 1 https://github.com/canonical/go-dqlite.git

RUN apt-get -y install sqlite3 libsqlite3-dev && \
    add-apt-repository -y ppa:dqlite/stable && \
    apt-get -y install dqlite libdqlite-dev && \
    apt-get -y autoremove && \
    rm -rf /var/lib/apt/lists/*

ARG MODULES_BIN

ENV PATH=${PATH}:/usr/local/go/bin
# WORKDIR /dataplane-manager
