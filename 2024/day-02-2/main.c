#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void levels_dump(int levels[], int valid) {
  int count = sizeof(*levels);
  printf("Levels: ");
  for (int i = 0; i < 16; i++) {
    printf("%d ", levels[i]);
  }

  printf("Valid: %d\n", valid);
}

void levels_copy(int levels[], int copy[], int skip_pos) {
  for (int i = 0, pos = 0; levels[i] != 0; i++) {
    if (i == skip_pos) {
      continue;
    }
    copy[pos] = levels[i];
    pos++;
  }
}

int levels_check(int levels[]) {
  int asc = 1;
  if (levels[0] > levels[1]) {
    asc = 0;
  }

  for (size_t i = 0; i < 16; i++) {
    if (levels[i] == 0 || levels[i + 1] == 0) {
      break;
    }

    if (asc == 1 && levels[i] > levels[i + 1]) {
      return 0;
    }
    if (asc == 0 && levels[i] < levels[i + 1]) {
      return 0;
    }

    int diff = abs(levels[i] - levels[i + 1]);
    if (diff < 1 || diff > 3) {
      return 0;
    }
  }

  return 1;
}

int main(void) {
  int total = 0;
  char line[256];

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  while (fgets(line, sizeof(line), file)) {
    int levels[16] = {};
    char *saveptr = NULL;
    char *token = strtok_r(line, " ", &saveptr);

    int i = 0;
    while (token != NULL) {
      levels[i] = atoi(token);
      token = strtok_r(NULL, " ", &saveptr);
      i++;
    }

    if (levels_check(levels) == 1) {
      total++;
    } else {
      for (int skip_pos = 0; levels[skip_pos] != 0; skip_pos++) {
        int copy[16] = {};
        levels_copy(levels, copy, skip_pos);
        if (levels_check(copy) == 1) {
          total++;
          break;
        };
      }
    }
  }

  pclose(file);

  printf("total: %d\n", total);
}
