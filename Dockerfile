FROM golang:1.22.4-bullseye

WORKDIR /meguca

RUN apt-get update && apt-get install -y \
  pkg-config \
  libopencv-dev \
  libgeoip-dev \
  libwebp-dev \
  ffmpeg \
  libavcodec-dev \
  libavformat-dev \
  libswscale-dev \
  build-essential \
  postgresql-client \
  && apt-get clean

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make all

EXPOSE 8000
