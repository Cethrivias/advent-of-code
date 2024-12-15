#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Ids {
  int len;
  int vals[1024];
};

void ids_add(struct Ids *ids, int id) {
  ids->vals[ids->len] = id;
  ids->len++;
}

void ids_sort(struct Ids *ids) {
  for (int i = 0; i < ids->len - 1; i++) {
    for (int j = i + 1; j < ids->len; j++) {
      int tmp = ids->vals[j];
      if (ids->vals[i] > ids->vals[j]) {
        ids->vals[j] = ids->vals[i];
        ids->vals[i] = tmp;
      }
    }
  }
}

void ids_print(struct Ids *ids) {
  printf("Ids->vals: ");
  for (int i = 0; i < ids->len; i++) {
    printf("%d", ids->vals[i]);
  }
  printf("\n");
}

int parse_location_id(char line[], size_t size) {
  char num[8] = "";

  strncpy(num, line, size);
  if (strlen(num) != size) {
    perror("Error while parsing the id. Copied number size is invalid\n");
    exit(EXIT_FAILURE);
  }

  return atoi(num);
}

int main(void) {
  int total = 0;
  char line[256];
  struct Ids left = {.len = 0, .vals = {}};
  struct Ids right = {.len = 0, .vals = {}};

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  while (fgets(line, sizeof(line), file)) {
    printf("Line: %s", line);

    int len = strlen(line);
    int l = 0, r = 0;
    int idx = 0;

    do {
      idx++;
    } while (line[idx] != ' ');

    l = parse_location_id(line, idx);

    do {
      idx++;
    } while (line[idx] == ' ');

    r = parse_location_id(&line[idx], len - idx - 1);
    printf("left: %d, right: %d\n", l, r);

    ids_add(&left, l);
    ids_add(&right, r);
  }

  ids_sort(&left);
  ids_sort(&right);
  ids_print(&left);
  ids_print(&right);

  for (int i = 0; i < left.len; i++) {
    total += abs(left.vals[i] - right.vals[i]);
  }

  pclose(file);

  printf("total: %d\n", total);
}
