#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#define ARR_SIZE 130
#define INPUT "input"

struct Position {
  int row;
  int col;
};

int total = 0;
struct Position directions[4] = {{-1, 0}, {0, 1}, {1, 0}, {0, -1}};
int direction = 0;
char line[256];
char map[ARR_SIZE][ARR_SIZE] = {};
int route[ARR_SIZE][ARR_SIZE] = {};

int main(void) {

  FILE *file = fopen(INPUT, "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  struct Position start = {};
  for (size_t row = 0; fgets(line, sizeof(line), file); row++) {
    memcpy(map[row], line, ARR_SIZE);

    if (start.row == 0 && start.col == 0) {
      for (size_t i = 0; i < ARR_SIZE; i++) {
        if (map[row][i] == '^') {
          start.row = row;
          start.col = i;
        }
      }
    }
  }

  pclose(file);

  // Finding normal route.
  // It does not make sense to put obstacles if guard is not going to hit them
  struct Position guard = {start.row, start.col};
  while (1) {
    route[guard.row][guard.col] = 1;
    struct Position next = {guard.row + directions[direction].row,
                            guard.col + directions[direction].col};

    // out of bounds
    if (next.row < 0 || next.row >= ARR_SIZE || next.col < 0 ||
        next.col >= ARR_SIZE) {
      break;
    }

    // Hit some obstacle -> turn
    if (map[next.row][next.col] == '#') {
      direction++;
      if (direction == 4) {
        direction = 0;
      }

      continue;
    }

    guard.row = next.row;
    guard.col = next.col;
  }

  for (size_t i = 0; i < ARR_SIZE; i++) {
    for (size_t j = 0; j < ARR_SIZE; j++) {
      if (!route[i][j]) {
        continue;
      }

      int map_visited[ARR_SIZE][ARR_SIZE] = {};
      struct Position guard = {start.row, start.col};
      direction = 0;

      map_visited[guard.row][guard.col] |= (1 << direction);
      map[i][j] = '#';

      while (1) {
        struct Position next = {guard.row + directions[direction].row,
                                guard.col + directions[direction].col};

        // out of bounds
        if (next.row < 0 || next.row >= ARR_SIZE || next.col < 0 ||
            next.col >= ARR_SIZE) {
          break;
        }

        // Next cell was already visited with the same direction -> loop
        // Inc total and try next
        if (map_visited[next.row][next.col] & (1 << direction)) {
          total++;
          break;
        }

        // Hit some obstacle -> turn
        if (map[next.row][next.col] == '#') {
          direction++;
          if (direction == 4) {
            direction = 0;
          }

          continue;
        }

        guard.row = next.row;
        guard.col = next.col;
        map_visited[guard.row][guard.col] |= (1 << direction);
      }

      map[i][j] = '.';
    }
  }

  printf("total: %d\n", total);
}
