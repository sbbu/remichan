FROM golang:1.22.4-bullseye

WORKDIR /meguca

# Install dependencies
RUN apt-get update && apt-get install -y \
  pkg-config \
  libgeoip-dev \
  libwebp-dev \
  ffmpeg \
  libavcodec-dev \
  libavformat-dev \
  libswscale-dev \
  build-essential \
  cmake \
  python3-dev \
  python3-numpy \
  curl \
  && curl -fsSL https://deb.nodesource.com/setup_16.x | bash - \
  && apt-get install -y nodejs \
  && apt-get clean

# Install OpenCV properly with architecture-aware paths
RUN apt-get update && apt-get install -y \
  libopencv-dev \
  && mkdir -p /usr/lib/$(uname -m)-linux-gnu/pkgconfig \
  && mkdir -p /meguca/www/videos \
  && ln -s $(find /usr/lib -name "opencv4.pc") /usr/lib/$(uname -m)-linux-gnu/pkgconfig/opencv.pc

# Set pkg-config path dynamically
ENV PKG_CONFIG_PATH="/usr/lib/$(uname -m)-linux-gnu/pkgconfig:/usr/local/lib/pkgconfig:${PKG_CONFIG_PATH}"

COPY go.mod go.sum ./
RUN go mod download

COPY package*.json ./
RUN npm install

COPY . .
RUN make all

EXPOSE 8000
CMD ["/meguca/meguca", "-a", "0.0.0.0:8000"]
