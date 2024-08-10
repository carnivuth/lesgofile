#!/bin/bash
useradd -m -b /var/lib lesgofile
mkdir -p /etc/lesgofile
chown -R lesgofile:lesgofile /etc/lesgofile
