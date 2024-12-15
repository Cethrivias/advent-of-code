#include <ctype.h>
#include <limits.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>

int try(uint64_t target, uint64_t numbers[],
        size_t numbers_size, size_t numbers_pos, uint64_t current) {
  if (numbers_pos == numbers_size) {
    return current == target;
  }
  if (current > target) {
    return 0;
  }

  if (try(target, numbers, numbers_size, numbers_pos + 1,
          current * numbers[numbers_pos])) {
    return 1;
  }

  if (try(target, numbers, numbers_size, numbers_pos + 1,
          current + numbers[numbers_pos])) {
    return 1;
  }

  return 0;
}

int main(int argc, char **argv) {
  uint64_t total = 0;
  char line[256];

  if (argc < 2) {
    perror("Input filename is required");
    exit(1);
  }
  char *filename = argv[1];

  FILE *file = fopen(filename, "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  uint64_t input[32] = {};
  while (fgets(line, sizeof(line), file)) {

    char num[32] = {};
    size_t num_pos = 0;
    size_t line_pos = 0;
    size_t col = 0;

    for (size_t line_pos = 0; line[line_pos] != '\0'; line_pos++) {
      if (isdigit(line[line_pos])) {
        num[num_pos] = line[line_pos];
        num_pos++;

        continue;
      }

      if (num_pos != 0) {
        input[col] = strtoull(num, NULL, 10);
        col++;
        memset(num, 0, num_pos);
        num_pos = 0;
      }
    }

    char signs[32] = {};
    if (try(input[0], &input[1], col - 1, 1, input[1])) {
      total += input[0];
    }
  }

  pclose(file);

  printf("total: %llu\n", total);
}
