#include <cuda_runtime.h>
#include <stdio.h>
#include <stdlib.h>

#define STEPS 10000
#define BLOCK_SIZE 16
#define IDX(x, y, L) ((y) * L + (x))

void random_init(int *grid, int size) {
  for (int i = 0; i < size * size; ++i) {
    grid[i] = (float)rand() / RAND_MAX < 0.5 ? 0 : 1;
  }
}

__global__ void evolveKernel(int *curr, int *next, int size) {
  int x = blockIdx.x * blockDim.x + threadIdx.x;
  int y = blockIdx.y * blockDim.y + threadIdx.y;

  if (x >= size || y >= size) {
    return;
  }

  int count = 0;
  for (int i = -1; i <= 1; i++) {
    for (int j = -1; j <= 1; j++) {
      if (i == 0 && j == 0) {
        continue;
      }
      int nb_idx = IDX(x + j, y + i, size);
      if (nb_idx > 0 && nb_idx < size * size) {
        count += curr[nb_idx];
      }
    }
  }

  int cell = curr[IDX(x, y, size)];
  if (cell == 1) {
    next[IDX(x, y, size)] = (count == 2 || count == 3) ? 1 : 0;
  } else {
    next[IDX(x, y, size)] = (count == 3) ? 1 : 0;
  }
}

int main(int argc, char **argv) {
  srand(2001);

  int size = atoi(argv[1]);

  int *h_grid = (int *)malloc(size * size * sizeof(int));
  int *h_temp = (int *)malloc(size * size * sizeof(int));

  random_init(h_grid, size);

  int *d_grid, *d_temp;
  cudaMalloc((void **)&d_grid, size * size * sizeof(int));
  cudaMalloc((void **)&d_temp, size * size * sizeof(int));

  cudaMemcpy(d_grid, h_grid, size * size * sizeof(int), cudaMemcpyHostToDevice);

  dim3 dimBlock(BLOCK_SIZE, BLOCK_SIZE);
  dim3 dimGrid((size + dimBlock.x - 1) / dimBlock.x,
               (size + dimBlock.y - 1) / dimBlock.y);

  int steps = STEPS;
  while (steps--) {
    evolveKernel<<<dimGrid, dimBlock>>>(d_grid, d_temp, size);

    int *swap = d_grid;
    d_grid = d_temp;
    d_temp = swap;
  }

  cudaMemcpy(h_grid, d_grid, size * size * sizeof(int), cudaMemcpyDeviceToHost);

  // for (int i = 0; i < size * size; ++i) {
  //   printf("%d ", h_grid[i]);
  // }
  // printf("\n");

  cudaFree(d_grid);
  cudaFree(d_temp);

  free(h_grid);
  free(h_temp);

  return 0;
}
