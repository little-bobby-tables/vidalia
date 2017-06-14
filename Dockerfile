FROM debian:jessie-backports

RUN apt-get update \
 && apt-get -y install git curl build-essential pkg-config \
libjpeg-turbo-progs libpng-dev libdjvulibre-dev \
libavformat-dev libmpg123-dev libsamplerate-dev libsndfile-dev \
cimg-dev libavcodec-dev libswscale-dev ffmpeg \
libmagic-dev

RUN curl -O https://www.imagemagick.org/download/ImageMagick.tar.gz \
 && tar xzf ImageMagick.tar.gz \
 && cd ImageMagick* \
 && ./configure \
 && make \
 && make install

RUN curl -O https://storage.googleapis.com/golang/go1.8.3.linux-amd64.tar.gz \
 && tar xf go1.8.3.linux-amd64.tar.gz \
 && mv go /usr/local \
 && echo 'export GOPATH=/go' >> ~/.bashrc \
 && echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> ~/.bashrc

RUN curl -Lo phash.tar.gz https://github.com/hszcg/pHash-0.9.6/tarball/master \
 && tar xzf phash.tar.gz \
 && cd hszcg-pHash-0.9.6-0548356/pHash-0.9.6 \
 && ./configure \
 && make \
 && make install \
 && ldconfig /usr/local/lib

RUN bash -c "source ~/.bashrc; \
go get gopkg.in/gographics/imagick.v3/imagick; \
go get github.com/kavu/go-phash; \
go get github.com/rakyll/magicmime; \
go get github.com/stretchr/testify;"

ENV VIDALIA_PATH /go/src/vidalia/test/hierarchy
