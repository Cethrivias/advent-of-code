#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define PAGES_SIZE 100

struct PageRules {
  size_t len;
  int pages[PAGES_SIZE];
};

struct PageRules page_rules_new() {
  struct PageRules page_rules = {.len = 0, .pages = {}};
  return page_rules;
}

void pages_rules_add(struct PageRules pages_rules[], int curr_page,
                     int prev_page) {
  struct PageRules *page_rules = &pages_rules[curr_page];
  page_rules->pages[page_rules->len] = prev_page;
  page_rules->len++;
}

int pages_rules_check(struct PageRules pages_rules[], int curr_page,
                      int prev_pages[]) {
  for (size_t i = 0; prev_pages[i] != 0; i++) {
    struct PageRules *page_rules = &pages_rules[prev_pages[i]];

    for (size_t j = 0; j <= page_rules->len; j++) {
      if (page_rules->pages[j] == curr_page) {
        return 0;
      }
    }
  }

  return 1;
}

void pages_rules_dump(struct PageRules pages_rules[]) {
  for (size_t i = 0; i < PAGES_SIZE; i++) {
    if (pages_rules[i].len == 0) {
      continue;
    }
    printf("Page: %lu ->", i);
    for (size_t j = 0; j < pages_rules[i].len; j++) {
      printf(" %d", pages_rules[i].pages[j]);
    }
    printf("\n");
  }
}

int main(void) {
  int total = 0;
  char line[256];

  struct PageRules pages_rules[PAGES_SIZE];
  for (size_t i = 0; i < PAGES_SIZE; i++) {
    pages_rules[i] = page_rules_new();
  }

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
    int prev_page = atoi(num);

    num[0] = line[3];
    num[1] = line[4];
    int curr_page = atoi(num);

    pages_rules_add(pages_rules, curr_page, prev_page);
  }

  while (fgets(line, sizeof(line), file)) {
    char *start;
    int valid = 1;
    int pages[32] = {};
    int pages_len = 0;
    start = &line[0];
    while (start[0] != '\0') {
      num[0] = start[0];
      num[1] = start[1];
      pages[pages_len] = atoi(num);
      if (!pages_rules_check(pages_rules, pages[pages_len], pages)) {
        valid = 0;
        break;
      }
      start = &start[3];
      pages_len++;
    }

    if (valid) {
      total += pages[pages_len / 2];
    }
  }

  pclose(file);

  printf("total: %d\n", total);
}
