#include <stdio.h>
#include <stdlib.h>

#define STEPS 10000
#define IDX(x, y, L) ((y) * L + (x))

void random_init(int *grid, int cells) {
  for (int i = 0; i < cells; ++i) {
    grid[i] = (float)rand() / RAND_MAX < 0.5 ? 0 : 1;
  }
}

void glider_init(int *grid, int size) {
  for (int i = 0; i < size * size; ++i) {
    grid[i] = 0;
  }
  grid[IDX(size / 2, size / 2, size)] = 1;
  grid[IDX(size / 2, (size / 2 - 1), size)] = 1;
  grid[IDX((size / 2 + 1), (size / 2 - 1), size)] = 1;
  grid[IDX((size / 2 - 1), size / 2, size)] = 1;
  grid[IDX(size / 2, (size / 2 + 1), size)] = 1;
}

int count_alive(int *grid, int size, int x, int y) {
  int count = 0;
  for (int i = -1; i <= 1; i++) {
    for (int j = -1; j <= 1; j++) {
      if (i == 0 && j == 0) {
        continue;
      }
      int nb_idx = IDX(x + j, y + i, size);
      if (nb_idx > 0 && nb_idx < size * size) {
        count += grid[nb_idx];
      }
    }
  }
  return count;
}

void evolve(int *curr, int *next, int size) {
  for (int y = 0; y < size; y++) {
    for (int x = 0; x < size; x++) {
      int count = count_alive(curr, size, x, y);
      int cell = curr[IDX(x, y, size)];
      if (cell == 1) {
        next[IDX(x, y, size)] = (count == 2 || count == 3) ? 1 : 0;
      } else {
        next[IDX(x, y, size)] = (count == 3) ? 1 : 0;
      }
    }
  }
}

int main(int argc, char **argv) {
  srand(2001);

  int size = atoi(argv[1]);

  int *grid = (int *)malloc(size * size * sizeof(int));
  int *temp = (int *)malloc(size * size * sizeof(int));

  // random_init(grid, size);
  glider_init(grid, size);

  int steps = STEPS;
  while (steps--) {
    evolve(grid, temp, size);
    for (int j = 0; j < size; j++) {
      for (int i = 0; i < size; i++) {
        grid[IDX(i, j, size)] = temp[IDX(i, j, size)];
      }
    }
    for (int i = 0; i < size * size; ++i)
      printf("%d ", grid[i]);
    printf("\n");
  }

  // for (int i = 0; i < size * size; ++i)
  //   printf("%d ", grid[i]);
  // printf("\n");

  free(grid);
  free(temp);

  return 0;
}
