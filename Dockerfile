FROM		debian:wheezy
MAINTAINER	Cedric Bosdonnat <cbosdonnat@suse.com>

# Deps
RUN echo 'deb http://http.debian.net/debian wheezy-backports main' >> /etc/apt/sources.list.d/backports.list
RUN apt-get update && \
    DEBCONF_FRONTEND=noninteractive DEBIAN_FRONTEND=noninteractive apt-get install -y \
    curl \
    gcc \
	libvirt0=1.2.* \
	libvirt-dev=1.2.* \
	libvirt-bin=1.2.*

# Install Go
RUN curl -sSL https://golang.org/dl/go1.3.1.src.tar.gz | tar -v -C /usr/local -xz
ENV GOROOT	/usr/local/go
ENV PATH    /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH  /go:/go/src/github.com/docker/docker/vendor
RUN cd /usr/local/go/src && ./make.bash --no-clean 2>&1

WORKDIR	/libvirt-go
COPY . /libvirt-go

ENTRYPOINT ["./run-libvirt"]
