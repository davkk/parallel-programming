#!/usr/bin/env bash

javac -d . src/main/java/com/github/davkk/Mandelbrot.java
javac -d . src/main/java/com/github/davkk/Benchmark.java

parallel -j 1 --progress \
    java com.github.davkk.Benchmark {1} {2} \>\> output/data/{1}.dat \
    ::: single threads pool-once pool-every pool-blocks-less pool-blocks-more \
    ::: 16 32 64 128 256 512 1024 2048 4096 8192

fd . output/data/ | uv run python plot.py
