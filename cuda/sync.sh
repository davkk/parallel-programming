#!/usr/bin/env zsh

jupytext cuda.py --to ipynb
# rclone copy -v ./main.cu GoogleDrive:/cuda --exclude "/.venv/**"
rclone bisync --resync --verbose --exclude "/.venv/**" --exclude "*.out" . GoogleDrive:/cuda
