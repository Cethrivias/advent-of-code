#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define PAGES_SIZE 100

struct SizedArray {
  size_t size;
  int values[PAGES_SIZE];
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

struct PageRules {
  int exists;
  struct SizedArray req_pages;
};

void page_rules_add_req_page(struct PageRules *page_rules, int req_page) {
  page_rules->exists = 1;
  append(&page_rules->req_pages, req_page);
}

int main(void) {
  int total = 0;
  char line[256];

  struct PageRules pages_rules[PAGES_SIZE] = {};

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  char num[3];
  while (fgets(line, sizeof(line), file)) {
    if (line[0] == '\n') {
      break;
    }

    num[0] = line[0];
    num[1] = line[1];
    int req_page = atoi(num);

    num[0] = line[3];
    num[1] = line[4];
    int page = atoi(num);

    page_rules_add_req_page(&pages_rules[page], req_page);
  }

  while (fgets(line, sizeof(line), file)) {
    char *start;
    struct SizedArray pages = {};
    start = &line[0];
    while (start[0] != '\0') {
      num[0] = start[0];
      num[1] = start[1];
      append(&pages, atoi(num));
      start = &start[3];
    }

    int pos = 0;
    int valid = 1;
    for (size_t i = 0; i < pages.size; i++) {
      struct SizedArray *req_pages = &pages_rules[pages.values[i]].req_pages;
      for (size_t j = i + 1; j < pages.size; j++) {
        if (contains(req_pages, pages.values[j])) {
          swap(&pages, i, j);
          valid = 0;
          i--;
          break;
        }
      }
    }

    if (!valid) {
      total += pages.values[pages.size / 2];
    }
  }

  pclose(file);

  printf("total: %d\n", total);
}
