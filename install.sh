#!/usr/bin/env bash

HOST_TARGET=${HOST_TARGET:-pi@tina}

# build
GOARM=6 GOARCH=arm GOOS=linux go build -o txt-ink-armv6

# copy
rsync -zarvh txt-ink-armv6 image.py ${HOST_TARGET}:~/

# TODO: write the unit files
# TODO: systemctl daemon-reload
# TODO: systemctl enable
