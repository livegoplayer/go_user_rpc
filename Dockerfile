#
# NOTE: THIS DOCKERFILE IS GENERATED VIA "update.sh"
#
# PLEASE DO NOT EDIT IT DIRECTLY.
#
FROM golang:latest

RUN sed -i s@/deb.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list
RUN sed -i s@/security.debian.org/@/mirrors.aliyun.com/@g /etc/apt/sources.list

# 切换到非交互模式避免警告
ENV DEBIAN_FRONTEND=noninteractive
# Locale
ENV LOCALE zh_CN.UTF-8
ENV TZ=Asia/Shanghai

RUN apt-get -qq update && \
  apt-get -yqq install \
  make \ 
  build-essential \
  apt-utils \
  wget \
  apt-transport-https lsb-release ca-certificates

# 常用工具 Install some basic tools needed for deployment
RUN apt-get -qq update && \
  apt-get -yqq install \
  debconf-utils \
  debconf \
  default-mysql-client \
  locate \
  locales \
  curl \
  unzip \
  patch \
  rsync \
  vim \
  nano \
  htop \
  tzdata \
  ntpdate \ 
  openssh-client \
  git \
  bash-completion \
  locales \
  libjpeg-turbo-progs libjpeg-progs \
  pngcrush optipng \
  perl  \
  python-pip \
  && apt-get -q autoclean && rm -rf /var/lib/apt/lists/*

# Install locale
RUN apt-get -qq update && \
  sed -i -e "s/# $LOCALE/$LOCALE/" /etc/locale.gen && \
  echo "LANG=$LOCALE">/etc/default/locale && \
  dpkg-reconfigure --frontend=noninteractive locales && \
  update-locale LANG=$LOCALE && \
  apt-get -q autoclean && rm -rf /var/lib/apt/lists/*

#go env
ENV GOPROXY=https://goproxy.io
ENV GO111MODULE=on
