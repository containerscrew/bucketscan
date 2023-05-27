#!/usr/bin/env bash

BUCKETSCAN_LATEST_VERSION=$(curl -s https://api.github.com/repos/containerscrew/bucketscan/releases/latest | jq -r ".name")
BUCKETSCAN_CLI_ARCH=amd64
if [ "$(uname -m)" = "aarch64" ]; then BUCKETSCAN_CLI_ARCH=arm64; fi
curl -L --fail --remote-name-all https://github.com/containerscrew/bucketscan/releases/download/${BUCKETSCAN_LATEST_VERSION}/bucketscan-linux-${BUCKETSCAN_CLI_ARCH}.deb
sudo dpkg -i bucketscan-linux-${BUCKETSCAN_CLI_ARCH}.deb
rm dpkg -i bucketscan-linux-${BUCKETSCAN_CLI_ARCH}.deb
