import sys

import numpy as np
from matplotlib import pyplot as plt

size = int(sys.argv[1])


def to_grid(line: str):
    grid = np.fromstring(line, sep=" ").T
    grid = grid.reshape(size, size)
    return grid


first = to_grid(sys.stdin.readline())

plt.ion()
fig, ax = plt.subplots(figsize=(8, 8))
im = ax.imshow(first, cmap="binary", interpolation="nearest")
ax.set_xticks([], [])
ax.set_yticks([], [])

try:
    while line := sys.stdin.readline():
        im.set_data(to_grid(line))
        fig.canvas.flush_events()
except KeyboardInterrupt:
    print("animation interrupt")

fig.tight_layout()
plt.ioff()
plt.show()
