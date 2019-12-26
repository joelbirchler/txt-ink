#!/usr/bin/env bash

HOST_TARGET=${HOST_TARGET:-pi@tina}
BUILD=${BUILD:-true}

# build
if [ "$BUILD" = "true" ]; then
  echo "Building txt-ink-armv6"
  GOARM=6 GOARCH=arm GOOS=linux go build -o txt-ink-armv6
fi

# copy
rsync -zarvh txt-ink-armv6 image.py enable.sh ${HOST_TARGET}:~/

# install and enable systemd units
ssh ${HOST_TARGET} "sudo /home/pi/enable.sh"
