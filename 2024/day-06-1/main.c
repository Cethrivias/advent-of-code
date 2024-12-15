#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ARR_SIZE 130
#define INPUT "input"

// 5330 - too low

struct SizedArray {
  size_t size;
  int values[ARR_SIZE];
};

void swap(struct SizedArray *arr, size_t i, size_t j) {
  int tmp = arr->values[j];
  arr->values[j] = arr->values[i];
  arr->values[i] = tmp;
}

int contains(struct SizedArray *arr, int el) {
  for (size_t i = 0; i < arr->size; i++) {
    if (arr->values[i] == el) {
      return 1;
    }
  }

  return 0;
}

int append(struct SizedArray *arr, int el) {
  arr->values[arr->size] = el;
  arr->size++;

  return 1;
}

int append_uniq(struct SizedArray *arr, int el) {
  if (contains(arr, el)) {
    return 0;
  }

  arr->values[arr->size] = el;
  arr->size++;

  return 1;
}

int pop(struct SizedArray *arr) {
  arr->size--;
  return arr->values[arr->size];
}

void sized_array_print(char prefix[], struct SizedArray *arr) {
  printf("%s", prefix);

  for (size_t i = 0; i < arr->size; i++) {
    printf("%d ", arr->values[i]);
  }

  printf("\n");
}

struct Position {
  int row;
  int col;
};

int within_bounds(int row, int col) {
  return row >= 0 && row < ARR_SIZE && col >= 0 && col < ARR_SIZE;
}

int main(void) {
  int total = 0;
  char line[256];
  char map[ARR_SIZE + 1][ARR_SIZE + 1] = {};
  struct Position directions[4] = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
  int direction = 0;
  struct Position guard = {};

  FILE *file = fopen(INPUT, "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  size_t row = 0;
  while (fgets(line, sizeof(line), file)) {
    // printf("Line: %s", line);
    memcpy(map[row], line, ARR_SIZE);

    if (guard.row == 0 && guard.col == 0) {
      for (size_t i = 0; i < ARR_SIZE; i++) {
        if (map[row][i] == '^') {
          guard.row = row;
          guard.col = i;
        }
      }
    }
    row++;
  }

  map[guard.row][guard.col] = 'x';
  total++;
  while (1) {
    int next_row = guard.row + directions[direction].row;
    int next_col = guard.col + directions[direction].col;

    if (!within_bounds(next_row, next_col)) {
      break;
    }

    // printf("Next: %c\n", map[next_row][next_col]);
    if (map[next_row][next_col] == '#') {
      direction++;
      if (direction == 4) {
        direction = 0;
      }
      continue;
    }

    if (map[next_row][next_col] != 'x') {
      total++;
      map[next_row][next_col] = 'x';
    }

    guard.row = next_row;
    guard.col = next_col;
  }

  pclose(file);

  printf("total: %d\n", total);
}
