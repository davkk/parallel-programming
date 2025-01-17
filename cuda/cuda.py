# ---
# jupyter:
#   jupytext:
#     text_representation:
#       extension: .py
#       format_name: percent
#       format_version: '1.3'
#       jupytext_version: 1.16.6
#   kernelspec:
#     display_name: Python 3
#     name: python3
# ---

# %%
from google.colab import drive
drive.mount("/content/drive", force_remount=True)

# %%
# !stat drive/MyDrive/cuda/main.cu
# !nvcc drive/MyDrive/cuda/main.cu
# !cat drive/MyDrive/cuda/sizes.txt | xargs -I{} bash -c "time ./a.out {}" 2> drive/MyDrive/cuda/output-cuda
