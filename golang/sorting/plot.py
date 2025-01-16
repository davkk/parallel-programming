import os
from pathlib import Path

import numpy as np
import scienceplots
from matplotlib import pyplot as plt

plt.style.use(["science", "ieee"])

methods = ["sequential", "parallel"]

OUTPUT = Path(__file__).parent / "output"

for method in methods:
    sizes = os.listdir(OUTPUT / method)
    means = [
        np.mean(
            [
                np.loadtxt(OUTPUT / method / size / run)[1]
                for run in os.listdir(OUTPUT / method / size)
                if run.endswith(".err")
            ]
        )
        for size in sizes
    ]
    plt.loglog(list(map(int, sizes)), means, ".-", label=method)

plt.title("Merge Sort - Parallel vs. Sequential")
plt.xlabel("Array Size")
plt.ylabel("Time [ns]")
plt.legend()

plt.tight_layout()
plt.savefig(Path(__file__).stem + ".pdf")
plt.show()
