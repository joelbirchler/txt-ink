#!/usr/bin/env bash

cat << UNITFILE > /etc/systemd/system/txt-ink-image.service
[Unit]
Description=Txt-Ink Image Writer
After=syslog.target

[Service]
Type=oneshot
WorkingDirectory=/home/pi
ExecStart=/home/pi/image.py message.png

[Install]
WantedBy=multi-user.target
UNITFILE


cat << UNITFILE > /etc/systemd/system/txt-ink-image.path
[Unit]
Description=Txt-Ink Image Writer

[Path]
PathModified=/home/pi/message.png

[Install]
WantedBy=multi-user.target
UNITFILE


cat << UNITFILE > /etc/systemd/system/txt-ink-server.service
[Unit]
Description=Txt-Ink Web Server
After=syslog.target network.target

[Service]
Type=simple
WorkingDirectory=/home/pi
ExecStart=/home/pi/txt-ink-armv6
Restart=on-failure

[Install]
WantedBy=multi-user.target
UNITFILE


systemctl daemon-reload

for unit in txt-ink-image.service txt-ink-image.path txt-ink-server.service; do
  systemctl enable $unit && systemctl start $unit
done
