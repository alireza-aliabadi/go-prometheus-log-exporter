#!/bin/bash

rsync -av \
--exclude venv \
--exclude .vscode \
--exclude .git \
--exclude .idea \
--exclude pkg \
--exclude *.*sh \
--exclude .pre-commit-config.yaml \
. $1:/root/monitoring