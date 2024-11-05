#!/bin/bash
set -e

# Tor daemon'ı başlat
service tor start

# Go uygulamasını başlat
exec /usr/local/bin/forum_monitoring_tools
