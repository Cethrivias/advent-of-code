#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define INPUT_SIZE 140

char input[INPUT_SIZE][INPUT_SIZE] = {};
char line[INPUT_SIZE + 2]; // + EOL and null terminator

struct Direction {
  int row;
  int col;
};
struct Direction new_direction(int row, int col) {
  struct Direction direction = {.row = row, .col = col};
  return direction;
}

int find_xmas(char input[][INPUT_SIZE], int a_row, int a_col) {
  int row, col;
  char other_char;

  struct Direction lines[2][2] = {// '\'
                                  {new_direction(-1, -1), new_direction(1, 1)},
                                  // '/'
                                  {new_direction(-1, 1), new_direction(1, -1)}};
  for (int i = 0; i < 2; i++) {
    other_char = '\0';
    for (int j = 0; j < 2; j++) {
      struct Direction *direction = &lines[i][j];
      row = a_row + direction->row;
      col = a_col + direction->col;

      if (row < 0 || row == INPUT_SIZE || col < 0 || col == INPUT_SIZE) {
        return 0;
      }

      if (input[row][col] != 'M' && input[row][col] != 'S') {
        return 0;
      }
      if (input[row][col] == other_char) {
        return 0;
      }

      other_char = input[row][col];
    }
  }

  return 1;
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
    memcpy(input[row], line, sizeof(input[row]));
    row++;
  }

  pclose(file);

  for (int row = 0; row < INPUT_SIZE; row++) {
    for (int col = 0; col < INPUT_SIZE; col++) {
      if (input[row][col] == 'A') {
        total += find_xmas(input, row, col);
      }
    }
  }

  printf("total: %d\n", total);
}
