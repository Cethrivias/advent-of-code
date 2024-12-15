#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT_SIZE 140

char input[INPUT_SIZE][INPUT_SIZE] = {};
char line[INPUT_SIZE + 2]; // + EOL and null terminator

char xmas[4] = "XMAS";

struct Direction {
  int row;
  int col;
};
struct Direction new_direction(int row, int col) {
  struct Direction direction = {.row = row, .col = col};
  return direction;
}

int find_xmas(char input[][INPUT_SIZE], int x_row, int x_col) {
  int total = 0;
  struct Direction directions[8] = {
      new_direction(-1, 0), new_direction(-1, 1), new_direction(0, 1),
      new_direction(1, 1),  new_direction(1, 0),  new_direction(1, -1),
      new_direction(0, -1), new_direction(-1, -1)};

  // printf("Found X (%c) at [%d:%d]\n", input[x_row][x_col], x_row, x_col);
  for (int i = 0; i < 8; i++) {
    int row = x_row;
    int col = x_col;
    struct Direction *direction = &directions[i];
    // printf("Going in [%d:%d]\n", direction->row, direction->col);

    for (int j = 1; j <= sizeof(xmas); j++) {
      if (j == sizeof(xmas)) {
        // printf("Found xmas [%d:%d]..[%d:%d]\n", x_row, x_col, row, col);
        total++;
        break;
      }

      row += direction->row;
      col += direction->col;
      if (row < 0 || row == INPUT_SIZE || col < 0 || col == INPUT_SIZE) {
        // printf("Out of bounds [%d:%d]\n", row, col);
        break;
      }
      if (input[row][col] != xmas[j]) {
        // printf("Not xmas\n");
        break;
      }
    }
  }

  return total;
}

int main(void) {
  int total = 0;

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  int row = 0;
  while (fgets(line, sizeof(line), file)) {
    printf("Line: %s", line);
    memcpy(input[row], line, sizeof(input[row]));
    row++;
  }

  pclose(file);

  for (int row = 0; row < INPUT_SIZE; row++) {
    for (int col = 0; col < INPUT_SIZE; col++) {
      if (input[row][col] == 'X') {
        total += find_xmas(input, row, col);
      }
    }
  }

  printf("total: %d\n", total);
}
