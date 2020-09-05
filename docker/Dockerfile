FROM golang:1.14

RUN apt-get update && \
  apt-get install -y \
  apt-utils \
  apt-transport-https \
  ca-certificates \
  curl \
  gnupg-agent

RUN curl -sSL https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ bionic main" > /etc/apt/sources.list.d/migrate.list
RUN apt-get update && \
  apt-get install -y migrate


WORKDIR /go/src/app

ENTRYPOINT ["docker/entrypoint.sh"]
