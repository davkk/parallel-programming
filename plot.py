import sys
from pathlib import Path

import numpy as np
from matplotlib import pyplot as plt

files = sys.stdin.read().splitlines()

for file in files:
    size, time = np.loadtxt(file).T
    plt.plot(
        size,
        time,
        "-",
        linewidth=1,
        color="lightgray",
    )
    plt.plot(
        size,
        time,
        "p",
        label=Path(file).stem,
        markersize=8,
        linewidth=4,
    )

plt.xscale("log", base=2)
plt.yscale("log")

plt.title("Benchmark - Mandelbrot", fontsize=14)
plt.xlabel("Size [px]", fontsize=12)
plt.ylabel("Time [s]", fontsize=12)

xticks = [16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192]
plt.xticks(xticks, labels=xticks)
plt.xlim(8, 16384)

plt.grid(linestyle="--", linewidth=0.5, alpha=0.3)

plt.legend()

plt.tight_layout()
plt.savefig("benchmark.pdf", bbox_inches="tight")
# plt.show()
